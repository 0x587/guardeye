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
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/eventpool"
	"github.com/0x587/guardeye/link/linkclient"
)

type IF interface {
	Run(ctx context.Context) IF
	Subscribe(topic string, handler func(transData []byte)) error
	Publish(topic string, transData []byte) error
	Call(topic string, reqTrans []byte) (rspTrans []byte, err error)
	GetServiceType(topic string) (*ServiceSchema, error)
	GetMessageType(topic string) (*MessageSchema, error)
	List() *linkclient.TypeListRsp
}

func New(ip string, port int, topicPatterns []string) IF {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", ip, port)}
	logx.Infof("foxglove connecting to %s", u.String())
	dialer := websocket.Dialer{
		Subprotocols: []string{"foxglove.websocket.v1"},
	}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		logx.Errorf("Failed to connect to WebSocket: %v", err)
	}

	res := &impl{
		conn:          conn,
		msgCh:         make(chan message, 1024),
		subscribers:   make(map[string][]func([]byte)),
		clientSendChs: make(map[string]*channelInfo),
		srvRspPool:    eventpool.New[uint32, []byte](),
		topicPatterns: topicPatterns,
	}
	return res
}

type impl struct {
	conn           *websocket.Conn
	connWriteMutex sync.Mutex
	serverSendChs  sync.Map // channelId: int -> *channelInfo

	clientSendChs      map[string]*channelInfo
	clientSendChsMutex sync.Mutex

	subscribers map[string][]func([]byte)
	msgCh       chan message

	srvChs sync.Map // serviceId: int -> *serviceInfo

	srvRspPool eventpool.IF[uint32, []byte]

	topicPatterns []string
}

func (i *impl) Run(ctx context.Context) IF {
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
				parallel.Map(i.subscribers[msg.Topic], func(f func([]byte), _ int) error {
					f(msg.TransData)
					return nil
				})
			}
		}
	}()
	go func() {
		for {
			type ServerSendMsg struct {
				Op string `json:"op"`
			}
			var msg ServerSendMsg
			select {
			case <-ctx.Done():
				return
			default:
				msgType, message, err := i.conn.ReadMessage()
				if err != nil {
					logx.Error(errors.Wrap(err, "foxglove: error reading websocket message"))
					return
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
	}()
	return i
}

func (i *impl) Subscribe(topic string, handler func(transData []byte)) error {
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

func (i *impl) Publish(topic string, cdrData []byte) (err error) {
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
			return errors.New(fmt.Sprintf("foxglove: topic %s not found", topic))
		}
		ch, err = i.clientAdvertise(ch)
		if err != nil {
			return err
		}
	}
	if err := i.clientMessageData(ch.Id, cdrData); err != nil {
		return err
	}
	return nil
}

func (i *impl) Call(topic string, reqCdr []byte) (rspCdr []byte, err error) {
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
		return nil, errors.New(fmt.Sprintf("foxglove: fail unknown service %s", topic))
	}
	callId, err := i.clientServiceCallReq(srv.Id, reqCdr)
	if err != nil {
		return nil, err
	}
	res, err := i.srvRspPool.Wait(ctx, callId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *impl) GetServiceType(topic string) (*ServiceSchema, error) {
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
		return nil, errors.New(fmt.Sprintf("foxglove: fail unknown service %s", topic))
	}
	res := &ServiceSchema{
		Name:           srv.Type,
		RequestSchema:  srv.RequestSchema,
		ResponseSchema: srv.ResponseSchema,
	}
	if res.RequestSchema == "" {
		res.RequestSchema = srv.Request.Schema
	}
	if res.ResponseSchema == "" {
		res.ResponseSchema = srv.Response.Schema
	}
	return res, nil
}

func (i *impl) GetMessageType(topic string) (*MessageSchema, error) {
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
		return nil, errors.New(fmt.Sprintf("foxglove: topic %s not found", topic))
	}
	return &MessageSchema{
		Schema: ch.Schema,
		Name:   ch.SchemaName,
	}, nil
}

func (i *impl) List() *linkclient.TypeListRsp {
	res := &linkclient.TypeListRsp{}
	i.serverSendChs.Range(func(key, value any) bool {
		channel := value.(*channelInfo)
		res.Messages = append(res.Messages, channel.Topic)
		return true
	})
	i.srvChs.Range(func(key, value any) bool {
		service := value.(*serviceInfo)
		res.Services = append(res.Services, service.Name)
		return true
	})
	return res
}

func (i *impl) handleServerAdvertise(msg []byte) {
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
		if !i.topicMatch(channel.Topic) {
			continue
		}
		logx.Infof("get channel %d %s", channel.Id, channel.Topic)
		i.serverSendChs.Store(channel.Id, &channelInfo{
			Id:         channel.Id,
			Topic:      channel.Topic,
			SchemaName: channel.SchemaName,
			Schema:     channel.Schema,
		})
	}
}

func (i *impl) handleServerUnAdvertise(msg []byte) {
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

func (i *impl) handleServerAdvertiseService(msg []byte) {
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
		if !i.topicMatch(service.Name) {
			continue
		}
		logx.Infof("get service %d %s", service.Id, service.Name)
		i.srvChs.Store(service.Id, &service)
	}
}

func (i *impl) handleServerUnAdvertiseService(msg []byte) {
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

func (i *impl) handleServerMessageData(msg []byte) {
	channelIdBuf := msg[1:5]
	channelId := binary.LittleEndian.Uint32(channelIdBuf)
	c, ok := i.serverSendChs.Load(int(channelId))
	if !ok {
		logx.Errorf("foxglove: channel not found, %d", channelId)
		return
	}
	select {
	case i.msgCh <- message{
		Topic:      c.(*channelInfo).Topic,
		SchemaName: c.(*channelInfo).SchemaName,
		TransData:  msg[13:],
	}:
	default:
		logx.Errorf("foxglove: output channel full")
		return
	}
}

func (i *impl) handleServerServiceRspData(msg []byte) {
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
	i.srvRspPool.Invoke(callID, payload)
}

func (i *impl) clientAdvertise(ch *channelInfo) (*channelInfo, error) {
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
	i.connWriteMutex.Lock()
	defer i.connWriteMutex.Unlock()
	err := i.conn.WriteJSON(msg)
	if err != nil {
		return nil, errors.Wrap(err, "foxglove: error writing advertise json message")
	}
	logx.Infof("client advertise %v", chi)
	return ch, nil
}

func (i *impl) clientMessageData(channelId int, data []byte) error {
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
	i.connWriteMutex.Lock()
	defer i.connWriteMutex.Unlock()
	err := i.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) clientServiceCallReq(serviceId int, data []byte) (uint32, error) {
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
	i.connWriteMutex.Lock()
	defer i.connWriteMutex.Unlock()
	err := i.conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		return 0, err
	}
	return callID, nil
}

func (i *impl) subscribe(ids ...int) error {
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
	i.connWriteMutex.Lock()
	defer i.connWriteMutex.Unlock()
	return errors.Wrap(i.conn.WriteJSON(msg), "foxglove: error writing subscribe json message")
}

func (i *impl) topicMatch(topic string) bool {
	for _, pattern := range i.topicPatterns {
		if _topicMatch(pattern, topic) {
			return true
		}
	}
	return false
}

func _topicMatch(pattern, topic string) bool {
	pParts := strings.Split(pattern, "/")
	tParts := strings.Split(topic, "/")

	for i := 0; i < len(pParts); i++ {
		if i >= len(tParts) {
			// 只有 # 可以匹配多余部分
			return pParts[i] == "#"
		}

		switch pParts[i] {
		case "#":
			// # 必须是最后一段
			return i == len(pParts)-1
		case "+":
			// + 匹配当前层级，继续匹配下一层
			continue
		default:
			// 精确匹配
			if pParts[i] != tParts[i] {
				return false
			}
		}
	}

	// 如果 topic 还有多余部分，且 pattern 没有 #，则不匹配
	return len(tParts) == len(pParts)
}
