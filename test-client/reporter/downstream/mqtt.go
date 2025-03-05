package downstream

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/downstream"
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
	if token := c.Subscribe(fmt.Sprintf("command/req/%s", cid.String()), 0, res.callback); token.Wait() && token.Error() != nil {
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

func (i *mqttImpl) callback(client mqtt.Client, msg mqtt.Message) {
	req := &downstream.CommandReq{}
	if err := json.Unmarshal(msg.Payload(), req); err != nil {
		logx.Error(err)
		return
	}
	go func() {
		rsp := &downstream.CommandRsp{
			Id: req.Id,
		}
		s := time.Now()
		res, err := ros.Do(req.Action, req.Data)
		logx.Infof("done %v", time.Now().Sub(s))
		if err != nil {
			logx.Error(err)
			rsp.Ok = false
			rsp.Data = err.Error()
		} else {
			rsp.Ok = true
			rsp.Data = res
		}
		rspBuf, err := json.Marshal(rsp)
		if err != nil {
			logx.Error(err)
			return
		}
		token := client.Publish(fmt.Sprintf("command/rsp/%s", i.cid), 0, false, rspBuf)
		if token.Wait(); token.Error() != nil {
			logx.Error(token.Error())
		}
		logx.Infof("reply %v", time.Now().Sub(s))
	}()
}
