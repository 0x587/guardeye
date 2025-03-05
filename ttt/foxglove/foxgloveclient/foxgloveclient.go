package foxgloveclient

import (
	"context"
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

	"github.com/0x587/guardeye/ttt/foxglove/foxgloveclient/roscdr"
)

type IF interface {
	Run(ctx context.Context) error
	Subscribe(topic string, handler func(jsonData string)) error
	Publish(topic, jsonData string) error
	Call(topic string, reqJson string) (rspJson string, err error)
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
	res := &Impl{
		conn: conn,
	}
	return res
}

type Impl struct {
	vm       *v8.Context
	conn     *websocket.Conn
	channels sync.Map // channelId: int -> channelInfo
}

func (i *Impl) Run(ctx context.Context) error {
	jsonMsgHandlers := map[string]func(msg []byte){
		"advertise":   i.handleServerAdvertise,
		"unadvertise": i.handleServerUnAdvertise,
	}
	binaryMsgHandlers := map[byte]func(msg []byte){
		1: i.handleServerMessageData,
	}
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

func (i *Impl) Subscribe(topic string, handler func(jsonData string)) error {
	//TODO implement me
	panic("implement me")
}

func (i *Impl) Publish(topic, jsonData string) error {
	//TODO implement me
	panic("implement me")
}

func (i *Impl) Call(topic string, reqJson string) (rspJson string, err error) {
	//TODO implement me
	panic("implement me")
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
	for _, channel := range advertiseMsg.Channels {
		logx.Infof("get channel %d", channel.Id)
		i.channels.Store(channel.Id, &channelInfo{
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
		return
	}
	for _, id := range unAdvertiseMsg.ChannelIds {
		i.channels.Delete(id)
	}
}

func (i *Impl) handleServerMessageData(msg []byte) {
	channelIdBuf := msg[1:5]
	channelId := binary.LittleEndian.Uint32(channelIdBuf)
	c, ok := i.channels.Load(int(channelId))
	if !ok {
		logx.Errorf("foxglove: channel not found, %d", channelId)
		return
	}
	schema := c.(*channelInfo).Schema
	res, err := roscdr.Parse(schema, msg[13:])
	if err != nil {
		logx.Errorf("foxglove: error reading message data: %v", err)
		return
	}

	select {
	case i.outputCh <- message{
		Topic:      c.(*channelInfo).Topic,
		SchemaName: c.(*channelInfo).SchemaName,
		JsonData:   res,
	}:
	default:
		logx.Errorf("foxglove: output channel full")
		return
	}
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
