package foxgloveclient

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/0x587/guardeye/common/eventpool"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
)

type IF interface {
	Run(ctx context.Context) error
	Subscribe(topic string, handler func(transData string)) error
	Publish(topic string, transData string) error
	Call(topic string, reqTrans string) (rspTrans string, err error)
}

func New(ip string, port int) IF {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", ip, port)}
	logx.Infof("foxglove connecting to %s", u.String())
	dialer := websocket.Dialer{
		Subprotocols: []string{"foxglove.websocket.v1"},
	}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		logx.Errorf("Failed to connect to WebSocket: %v", err)
	}

	client := lo.Must(grpc.NewClient("127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())))
	cli := foxgloveService.NewFoxgloveServiceClient(client)

	res := &Impl{
		conn:          conn,
		msgCh:         make(chan message, 1024),
		subscribers:   make(map[string][]func(string)),
		clientSendChs: make(map[string]*channelInfo),
		cli:           cli,
		srvRspPool:    eventpool.New[uint32, string](),
	}
	return res
}

type Impl struct {
	conn          *websocket.Conn
	serverSendChs sync.Map // channelId: int -> *channelInfo

	clientSendChs      map[string]*channelInfo
	clientSendChsMutex sync.Mutex

	subscribers map[string][]func(string)
	msgCh       chan message

	srvChs sync.Map // serviceId: int -> *serviceInfo

	srvRspPool eventpool.IF[uint32, string]

	cli foxgloveService.FoxgloveServiceClient
}

func (i *Impl) Run(ctx context.Context) error {
	jsonMsgHandlers := map[string]func(msg []byte){
		"advertise":         i.handleServerAdvertise,
		"unadvertise":       i.handleServerUnAdvertise,
		"advertiseServices": i.handleServerAdvertiseService,
	}
	binaryMsgHandlers := map[byte]func(msg []byte){
		0x01: i.handleServerMessageData,
		0x03: i.handleServerServiceRspData,
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-i.msgCh:
				parallel.Map(i.subscribers[msg.Topic], func(f func(string), _ int) error {
					f(msg.TransData)
					return nil
				})
			}
		}
	}()
	for {
		type ServerSendMsg struct {
			Op string `json:"op"`
		}
		var msg ServerSendMsg
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msgType, message, err := i.conn.ReadMessage()
			if err != nil {
				return errors.Wrap(err, "foxglove: error reading websocket message")
			}
			switch msgType {
			case websocket.TextMessage:
				if err := json.Unmarshal(message, &msg); err != nil {
					logx.Error(errors.Wrap(err, "foxglove: parse msg json fail"))
					continue
				}
				handleFunc := jsonMsgHandlers[msg.Op]
				if handleFunc != nil {
					handleFunc(message)
				}
			case websocket.BinaryMessage:
				opCode := message[0]
				handleFunc := binaryMsgHandlers[opCode]
				if handleFunc != nil {
					handleFunc(message)
				}
			}
		}
	}
}

func (i *Impl) Subscribe(topic string, handler func(transData string)) error {
	var ch *channelInfo
	i.serverSendChs.Range(func(key, value any) bool {
		channel := value.(*channelInfo)
		if channel.Topic == topic {
			ch = channel
			return false
		}
		return true
	})
	if ch == nil {
		return errors.New(fmt.Sprintf("topic %s not found", topic))
	}
	err := i.subscribe(ch.Id)
	if err != nil {
		return err
	}
	i.subscribers[topic] = append(i.subscribers[topic], handler)
	return nil
}

func (i *Impl) Publish(topic string, transData string) (err error) {
	i.clientSendChsMutex.Lock()
	ch, ok := i.clientSendChs[topic]
	i.clientSendChsMutex.Unlock()
	if !ok {
		i.serverSendChs.Range(func(key, value any) bool {
			channel := value.(*channelInfo)
			if channel.Topic == topic {
				ch = channel
				return false
			}
			return true
		})

		if ch == nil {
			return errors.New(fmt.Sprintf("topic %s not found", topic))
		}
		ch, err = i.clientAdvertise(ch)
		if err != nil {
			return err
		}
	}
	rsp, err := i.cli.CdrWrite(context.Background(), &foxgloveService.CdrWriteReq{
		RosSchema: ch.Schema,
		TransData: transData,
	})
	if err != nil {
		return err
	}
	if err := i.clientMessageData(ch.Id, rsp.Buf); err != nil {
		return err
	}
	return nil
}

func (i *Impl) Call(topic string, reqTrans string) (rspTrans string, err error) {
	ctx := context.Background()
	var srv *serviceInfo
	i.srvChs.Range(func(key, value any) bool {
		s := value.(*serviceInfo)
		if s.Name != topic {
			return true
		}
		srv = s
		return false
	})
	if srv == nil {
		return "", errors.Wrap(err, fmt.Sprintf("foxglove: fail unknown service %s", topic))
	}
	schema := srv.Request.Schema
	if schema == "" {
		schema = srv.RequestSchema
	}
	rsp, err := i.cli.CdrWrite(ctx, &foxgloveService.CdrWriteReq{
		RosSchema: schema,
		TransData: reqTrans,
	})
	if err != nil {
		return "", err
	}
	callId, err := i.clientServiceCallReq(srv.Id, rsp.Buf)
	if err != nil {
		return "", err
	}
	pollRes, err := i.srvRspPool.Wait(ctx, callId)
	if err != nil {
		return "", err
	}
	return pollRes, nil
}

func (i *Impl) handleServerAdvertise(msg []byte) {
	type AdvertiseMsg struct {
		Op       string        `json:"op"`
		Channels []channelInfo `json:"channels"`
	}
	var advertiseMsg AdvertiseMsg
	if err := json.Unmarshal(msg, &advertiseMsg); err != nil {
		logx.Error(err)
		return
	}
	for _, channel := range advertiseMsg.Channels {
		logx.Infof("get channel %d %s", channel.Id, channel.Topic)
		i.serverSendChs.Store(channel.Id, &channelInfo{
			Id:         channel.Id,
			Topic:      channel.Topic,
			SchemaName: channel.SchemaName,
			Schema:     channel.Schema,
		})
	}
}

func (i *Impl) handleServerUnAdvertise(msg []byte) {
	type UnAdvertiseMsg struct {
		Op         string `json:"op"`
		ChannelIds []int  `json:"channelIds"`
	}
	var unAdvertiseMsg UnAdvertiseMsg
	if err := json.Unmarshal(msg, &unAdvertiseMsg); err != nil {
		logx.Error(err)
		return
	}
	for _, id := range unAdvertiseMsg.ChannelIds {
		logx.Infof("del channel %d", id)
		i.serverSendChs.Delete(id)
	}
}

func (i *Impl) handleServerAdvertiseService(msg []byte) {
	type AdvertiseServiceMsg struct {
		Op       string        `json:"op"`
		Services []serviceInfo `json:"services"`
	}
	advertiseServiceMsg := &AdvertiseServiceMsg{}
	if err := json.Unmarshal(msg, advertiseServiceMsg); err != nil {
		logx.Error(err)
		return
	}
	for _, service := range advertiseServiceMsg.Services {
		logx.Infof("get service %d %s", service.Id, service.Name)
		i.srvChs.Store(service.Id, &service)
	}
}

func (i *Impl) handleServerUnAdvertiseService(msg []byte) {
	type UnAdvertiseServiceMsg struct {
		Op         string `json:"op"`
		ServiceIds []int  `json:"serviceIds"`
	}
	var unAdvertiseServiceMsg UnAdvertiseServiceMsg
	if err := json.Unmarshal(msg, &unAdvertiseServiceMsg); err != nil {
		logx.Error(err)
		return
	}
	for _, id := range unAdvertiseServiceMsg.ServiceIds {
		logx.Infof("del service %d", id)
		i.srvChs.Delete(id)
	}
}

func (i *Impl) handleServerMessageData(msg []byte) {
	channelIdBuf := msg[1:5]
	channelId := binary.LittleEndian.Uint32(channelIdBuf)
	c, ok := i.serverSendChs.Load(int(channelId))
	if !ok {
		logx.Errorf("foxglove: channel not found, %d", channelId)
		return
	}
	schema := c.(*channelInfo).Schema
	rsp, err := i.cli.CdrRead(context.Background(), &foxgloveService.CdrReadReq{
		RosSchema: schema,
		Buf:       msg[13:],
	})
	if err != nil {
		logx.Errorf("foxglove: error reading message data: %v", err)
		return
	}
	select {
	case i.msgCh <- message{
		Topic:      c.(*channelInfo).Topic,
		SchemaName: c.(*channelInfo).SchemaName,
		TransData:  rsp.TransData,
	}:
	default:
		logx.Errorf("foxglove: output channel full")
		return
	}
}

func (i *Impl) handleServerServiceRspData(msg []byte) {
	buf := bytes.NewBuffer(msg)
	var opcode byte
	if err := binary.Read(buf, binary.LittleEndian, &opcode); err != nil {
		logx.Error(err)
		return
	}
	var serviceID uint32
	if err := binary.Read(buf, binary.LittleEndian, &serviceID); err != nil {
		logx.Error(err)
		return
	}
	var callID uint32
	if err := binary.Read(buf, binary.LittleEndian, &callID); err != nil {
		logx.Error(err)
		return
	}
	var encodingLength uint32
	if err := binary.Read(buf, binary.LittleEndian, &encodingLength); err != nil {
		logx.Error(err)
		return
	}
	var encoding = make([]byte, encodingLength)
	if err := binary.Read(buf, binary.LittleEndian, &encoding); err != nil {
		logx.Error(err)
		return
	}
	payload, err := io.ReadAll(buf)
	if err != nil {
		logx.Error(err)
		return
	}
	s, ok := i.srvChs.Load(int(serviceID))
	if !ok {
		logx.Errorf("foxglove: service not found, %d", serviceID)
		return
	}
	schema := s.(*serviceInfo).Response.Schema
	if schema == "" {
		schema = s.(*serviceInfo).ResponseSchema
	}
	rsp, err := i.cli.CdrRead(context.Background(), &foxgloveService.CdrReadReq{
		RosSchema: schema,
		Buf:       payload,
	})
	if err != nil {
		logx.Errorf("foxglove: error reading message data: %v", err)
		return
	}
	i.srvRspPool.Invoke(callID, rsp.TransData)
}

func (i *Impl) clientAdvertise(ch *channelInfo) (*channelInfo, error) {
	type ClientAdvertiseMsg struct {
		Op       string        `json:"op"`
		Channels []channelInfo `json:"channels"`
	}
	i.clientSendChsMutex.Lock()
	defer i.clientSendChsMutex.Unlock()
	chi := channelInfo{
		Id:         len(i.clientSendChs) + 2,
		Topic:      ch.Topic,
		Encoding:   "cdr",
		Schema:     ch.Schema,
		SchemaName: ch.SchemaName,
	}
	msg := &ClientAdvertiseMsg{
		Op:       "advertise",
		Channels: []channelInfo{chi},
	}
	i.clientSendChs[chi.Topic] = &chi
	err := i.conn.WriteJSON(msg)
	if err != nil {
		return nil, errors.Wrap(err, "foxglove: error writing advertise json message")
	}
	logx.Infof("client advertise %v", chi)
	return ch, nil
}

func (i *Impl) clientMessageData(channelId int, data []byte) error {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, byte(0x01)); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, uint32(channelId)); err != nil {
		return err
	}
	if _, err := buf.Write(data); err != nil {
		return err
	}
	logx.Infof("send %v to channel %d", buf.Bytes(), channelId)
	err := i.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (i *Impl) clientServiceCallReq(serviceId int, data []byte) (uint32, error) {
	buf := bytes.NewBuffer(nil)
	if err := binary.Write(buf, binary.LittleEndian, byte(0x02)); err != nil {
		return 0, err
	}
	if err := binary.Write(buf, binary.LittleEndian, uint32(serviceId)); err != nil {
		return 0, err
	}
	callID := rand.Uint32()
	if err := binary.Write(buf, binary.LittleEndian, uint32(callID)); err != nil {
		return 0, err
	}
	encoding := "cdr"
	if err := binary.Write(buf, binary.LittleEndian, uint32(len(encoding))); err != nil {
		return 0, err
	}
	if err := binary.Write(buf, binary.LittleEndian, []byte(encoding)); err != nil {
		return 0, err
	}
	if _, err := buf.Write(data); err != nil {
		return 0, err
	}
	logx.Infof("send call: %d %v to service %d", callID, buf.Bytes(), serviceId)
	err := i.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return 0, err
	}
	return callID, nil
}

func (i *Impl) subscribe(ids ...int) error {
	type Subscription struct {
		Id        int `json:"id"`
		ChannelId int `json:"channelId"`
	}
	type ClientSubscribeMsg struct {
		Op            string         `json:"op"`
		Subscriptions []Subscription `json:"subscriptions"`
	}
	msg := ClientSubscribeMsg{
		Op: "subscribe",
		Subscriptions: lo.FilterMap(ids, func(id int, _ int) (Subscription, bool) {
			return Subscription{
				Id:        id,
				ChannelId: id,
			}, true
		}),
	}
	return errors.Wrap(i.conn.WriteJSON(msg), "foxglove: error writing subscribe json message")
}
