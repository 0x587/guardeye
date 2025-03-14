package logic

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"

	"github.com/0x587/guardeye/common/eventpool"
	"github.com/0x587/guardeye/common/polling"
	"github.com/0x587/guardeye/link/link"
)

type AgentConn interface {
	Call(ctx context.Context, cid string, req *link.LinkCommandDownstream) (rsp *link.LinkCommandUpstream, err error)
	List(ctx context.Context) (*link.AgentListRsp, error)
}

type mqttAgentConn struct {
	mqtt  mqtt.Client
	event eventpool.IF[string, *link.LinkCommandUpstream]
}

func newMqttAgentConnImpl(endpoint string) AgentConn {
	opts := mqtt.NewClientOptions().AddBroker(endpoint).
		SetClientID(fmt.Sprintf("service-api-%d", rand.Int())).
		SetKeepAlive(5 * time.Second).
		SetPingTimeout(1 * time.Second).
		SetAutoReconnect(true)
	mqttCli := mqtt.NewClient(opts)
	if token := mqttCli.Connect(); token.Wait() && token.Error() != nil {
		logx.Must(token.Error())
	}
	e := eventpool.New[string, *link.LinkCommandUpstream]()
	var cbf mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		rsp := &link.LinkCommandUpstream{}
		err := json.Unmarshal(message.Payload(), rsp)
		if err != nil {
			return
		}
		e.Invoke(rsp.Id, rsp)
	}
	token := mqttCli.Subscribe("command/rsp/#", 2, cbf)
	if token.Wait(); token.Error() != nil {
		logx.Must(token.Error())
	}
	return &mqttAgentConn{
		mqtt:  mqttCli,
		event: e,
	}
}

func (c *mqttAgentConn) Init(callback func(rsp *link.LinkCommandUpstream)) {
	var cbf mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		rsp := &link.LinkCommandUpstream{}
		err := json.Unmarshal(message.Payload(), rsp)
		if err != nil {
			return
		}
		c.event.Invoke(rsp.Id, rsp)
		callback(rsp)
	}
	token := c.mqtt.Subscribe("command/rsp/#", 2, cbf)
	if token.Wait(); token.Error() != nil {
		logx.Must(token.Error())
	}
}

func (c *mqttAgentConn) Send(cid string, req *link.LinkCommandDownstream) error {
	msgToAgentStr, err := json.Marshal(req)
	if err != nil {
		return err
	}
	token := c.mqtt.Publish(fmt.Sprintf("command/req/%s", cid), 2, false, msgToAgentStr)
	if token.Wait(); token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *mqttAgentConn) Call(ctx context.Context, cid string, req *link.LinkCommandDownstream) (rsp *link.LinkCommandUpstream, err error) {
	callID := uuid.New()
	req.Id = callID.String()
	if err := c.Send(cid, req); err != nil {
		return nil, err
	}
	rsp, err = c.event.Wait(ctx, callID.String())
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *mqttAgentConn) List(ctx context.Context) (*link.AgentListRsp, error) {
	//TODO implement me
	panic("implement me")
}

var linkRpcEvent = eventpool.New[string, *link.LinkCommandUpstream]()
var linkRpcPoll = polling.New[*link.LinkCommandDownstream, *link.LinkCommandUpstream](
	func(key string, rsp *link.LinkCommandUpstream) {
		linkRpcEvent.Invoke(rsp.GetId(), rsp)
	})

type rpcAgentConn struct {
	event eventpool.IF[string, *link.LinkCommandUpstream]
	poll  polling.IF[*link.LinkCommandDownstream, *link.LinkCommandUpstream]
}

func newRpcAgentConn() AgentConn {
	return &rpcAgentConn{
		event: linkRpcEvent,
		poll:  linkRpcPoll,
	}
}

func (c *rpcAgentConn) Call(ctx context.Context, cid string, req *link.LinkCommandDownstream) (*link.LinkCommandUpstream, error) {
	req.Id = uuid.New().String()
	if err := c.poll.Send(cid, req); err != nil {
		return nil, err
	}
	rsp, err := c.event.Wait(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *rpcAgentConn) List(ctx context.Context) (*link.AgentListRsp, error) {
	res := &link.AgentListRsp{Agents: make(map[string]string)}
	now := time.Now()
	for key, seeAt := range c.poll.List() {
		if seeAt.Add(time.Minute).After(now) {
			res.Agents[key] = seeAt.Format(time.RFC3339)
		}
	}
	return res, nil
}
