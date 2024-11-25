package wsclient

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	v8 "rogchap.com/v8go"
)

func New(ip string, port int, topics ...string) *Impl {
	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%d", ip, port)}
	logx.Infof("foxglove connecting to %s", u.String())
	dialer := websocket.Dialer{
		Subprotocols: []string{"foxglove.websocket.v1"},
	}
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		logx.Errorf("Failed to connect to WebSocket: %v", err)
	}
	res := &Impl{
		conn:       conn,
		channelIds: map[int]*channelInfo{},
		outputCh:   make(chan message, 100),
	}
	for _, topic := range topics {
		res.needSubscribeTopic.Store(topic, struct{}{})
	}
	return res
}

type Impl struct {
	vm                 *v8.Context
	conn               *websocket.Conn
	channelIds         map[int]*channelInfo
	channelMutex       sync.RWMutex
	outputCh           chan message
	needSubscribeTopic sync.Map
}

type channelInfo struct {
	Topic      string
	SchemaName string
	Schema     string
	reader     *foxgloveReader
}

type message struct {
	Topic      string
	SchemaName string
	Data       string
}

func (i *Impl) GetOutputChan() chan message {
	return i.outputCh
}

func (i *Impl) Handle() error {
	type ServerSendMsg struct {
		Op string `json:"op"`
	}
	var msg ServerSendMsg
	jsonMsgHandlers := map[string]func(msg []byte){
		"advertise":   i.handleServerAdvertise,
		"unadvertise": i.handleServerUnAdvertise,
	}
	binaryMsgHandlers := map[byte]func(msg []byte){
		1: i.handleServerMessageData,
	}
	for {
		_, message, err := i.conn.ReadMessage()
		if err != nil {
			return errors.Wrap(err, "foxglove: error reading websocket message")
		}

		if err := json.Unmarshal(message, &msg); err != nil {
			opCode := message[0]
			handleFunc := binaryMsgHandlers[opCode]
			if handleFunc != nil {
				handleFunc(message)
			}
			continue
		}

		handleFunc := jsonMsgHandlers[msg.Op]
		if handleFunc != nil {
			handleFunc(message)
		}
	}
}

func (i *Impl) handleServerAdvertise(msg []byte) {
	type AdvertiseMsg struct {
		Op       string `json:"op"`
		Channels []struct {
			Id         int    `json:"id"`
			Topic      string `json:"topic"`
			Encoding   string `json:"encoding"`
			SchemaName string `json:"schemaName"`
			Schema     string `json:"schema"`
		} `json:"channels"`
	}
	var advertiseMsg AdvertiseMsg
	if err := json.Unmarshal(msg, &advertiseMsg); err != nil {
		return
	}
	i.channelMutex.Lock()
	defer i.channelMutex.Unlock()
	var needSubscribeIds []int
	for _, channel := range advertiseMsg.Channels {
		i.channelIds[channel.Id] = &channelInfo{
			Topic:      channel.Topic,
			SchemaName: channel.SchemaName,
			Schema:     channel.Schema,
		}
		_, exist := i.needSubscribeTopic.Load(channel.Topic)
		if !exist {
			continue
		}
		needSubscribeIds = append(needSubscribeIds, channel.Id)
		i.needSubscribeTopic.Delete(channel.Topic)
	}
	if len(needSubscribeIds) > 0 {
		if err := i.subscribe(needSubscribeIds...); err != nil {
			logx.Errorf("foxglove: error subscribing channels: %v", err)
		}
	}
}

func (i *Impl) handleServerUnAdvertise(msg []byte) {
	type UnAdvertiseMsg struct {
		Op         string `json:"op"`
		ChannelIds []int  `json:"channelIds"`
	}
	var unAdvertiseMsg UnAdvertiseMsg
	if err := json.Unmarshal(msg, &unAdvertiseMsg); err != nil {
		return
	}
	i.channelMutex.Lock()
	defer i.channelMutex.Unlock()
	for _, id := range unAdvertiseMsg.ChannelIds {
		delete(i.channelIds, id)
	}
}

func (i *Impl) handleServerMessageData(msg []byte) {
	channelIdBuf := msg[1:5]
	channelId := binary.LittleEndian.Uint32(channelIdBuf)
	reader, err := i.getReader(int(channelId))
	if err != nil {
		logx.Errorf("foxglove: error getting reader: %v", err)
		return
	}
	res, err := reader.read(msg[13:])
	if err != nil {
		logx.Errorf("foxglove: error reading message data: %v", err)
		return
	}
	go func() {
		i.channelMutex.RLock()
		defer i.channelMutex.RUnlock()
		select {
		case i.outputCh <- message{
			Topic:      i.channelIds[int(channelId)].Topic,
			SchemaName: i.channelIds[int(channelId)].SchemaName,
			Data:       res,
		}:
		default:
			logx.Errorf("foxglove: output channel full")
			return
		}
	}()
}

func (i *Impl) getReader(channelId int) (*foxgloveReader, error) {
	i.channelMutex.RLock()
	c := i.channelIds[channelId]
	if c == nil {
		i.channelMutex.RUnlock()
		return nil, errors.New("foxglove: channel not found")
	}
	reader := c.reader
	if reader == nil {
		i.channelMutex.RUnlock()
		return i.createReader(channelId)
	}
	defer i.channelMutex.RUnlock()
	return reader, nil
}

func (i *Impl) createReader(channelId int) (*foxgloveReader, error) {
	i.channelMutex.Lock()
	defer i.channelMutex.Unlock()
	c := i.channelIds[channelId]
	if c == nil {
		return nil, errors.New("foxglove: channel not found")
	}
	vm, err := newJsVm()
	if err != nil {
		return nil, errors.Wrap(err, "foxglove: error creating new js vm")
	}
	reader, err := newFoxgloveReader(vm, c.Schema)
	if err != nil {
		return nil, errors.Wrap(err, "foxglove: error creating new foxglove reader")
	}
	c.reader = reader
	return reader, nil
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
			if i.channelIds[id] == nil {
				logx.Errorf("foxglove :Channel %d not found", id)
				return Subscription{}, false
			}
			return Subscription{
				Id:        id,
				ChannelId: id,
			}, true
		}),
	}
	return errors.Wrap(i.conn.WriteJSON(msg), "foxglove: error writing subscribe json message")
}
