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
		c:   c,
		cid: cid.String(),
	}

	// 订阅主题
	if token := c.Subscribe(fmt.Sprintf("command/req/%s", cid.String()), 2, res.callback); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return res, nil
}

type mqttImpl struct {
	c   mqtt.Client
	cid string
}

func (i *mqttImpl) Close() {
	i.c.Disconnect(250)
}

type commandReq struct {
	Action  string         `json:"action"`
	Id      string         `json:"id"`
	Payload map[string]any `json:"payload"`
}

type commandRsp struct {
	Id string `json:"id"`
	Ok bool   `json:"ok"`
}

func (i *mqttImpl) callback(client mqtt.Client, msg mqtt.Message) {
	req := &commandReq{}
	if err := json.Unmarshal(msg.Payload(), req); err != nil {
		logx.Error(err)
		return
	}
	if _, err := ros.Do(req.Action, req.Payload); err != nil {
		logx.Error(err)
		return
	}
	rsp := &commandRsp{
		Id: req.Id,
		Ok: true,
	}
	rspBuf, err := json.Marshal(rsp)
	if err != nil {
		logx.Error(err)
		return
	}
	token := client.Publish(fmt.Sprintf("command/rsp/%s", i.cid), 2, false, rspBuf)
	if token.Wait(); token.Error() != nil {
		logx.Error(token.Error())
	}
}
