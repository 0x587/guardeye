package downstream

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/test-client/reporter/ros"
)

type MqttCli interface {
	Close()
}

func NewMqtt(cid uuid.UUID) (MqttCli, error) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://emqxtcp.guardeye.shawnsiu.site:58701").
		SetClientID(fmt.Sprintf("agent_%s", cid.String()))

	opts.SetKeepAlive(60 * time.Second)
	// 设置消息回调处理函数
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	res := &mqttImpl{
		c: c,
	}

	// 订阅主题
	if token := c.Subscribe(fmt.Sprintf("command/%s", cid.String()), 2, res.callback); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return res, nil
}

type mqttImpl struct {
	c mqtt.Client
}

func (i *mqttImpl) Close() {
	i.c.Disconnect(250)
}

type msgPayload struct {
	Action  string         `json:"action"`
	Id      string         `json:"id"`
	Payload map[string]any `json:"payload"`
}

func (i *mqttImpl) callback(client mqtt.Client, msg mqtt.Message) {
	p := &msgPayload{}
	if err := json.Unmarshal(msg.Payload(), p); err != nil {
		logx.Error(err)
		return
	}
	if err := ros.Do(p.Action, p.Payload); err != nil {
		logx.Error(err)
		return
	}
}
