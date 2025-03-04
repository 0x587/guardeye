package downstream

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/downstream"
	"github.com/0x587/guardeye/common/eventpool"
	"github.com/0x587/guardeye/test-client/reporter/ros"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var mqttInitOnce sync.Once
var event = eventpool.New[string, *downstream.CommandRsp]()

func NewCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallLogic {
	mqttInitOnce.Do(func() {
		var callback mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
			rsp := &downstream.CommandRsp{}
			err := json.Unmarshal(message.Payload(), rsp)
			if err != nil {
				return
			}
			event.Invoke(rsp.Id, rsp)
		}
		token := svcCtx.Mqtt.Subscribe("command/rsp/#", 0, callback)
		if token.Wait(); token.Error() != nil {
			logx.Must(token.Error())
		}
	})
	return &CallLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallLogic) Call(req *types.DownstreamCallReq) (resp *types.DownstreamCallRsp, err error) {
	dataObj := make(map[string]any)
	if err := json.Unmarshal([]byte(req.Data), &dataObj); err != nil {
		return nil, err
	}
	topic, err := method2Topic(req.Method)
	if err != nil {
		return nil, err
	}
	pid := uuid.New()
	p := &downstream.CommandReq{
		Id:     pid.String(),
		Action: req.Action,
		Data: downstream.CommandReqData{
			RosTopic: topic,
			Data:     dataObj,
		},
	}
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	token := l.svcCtx.Mqtt.Publish(fmt.Sprintf("command/req/%s", req.Cid), 0, false, marshal)
	if token.Wait(); token.Error() != nil {
		return nil, token.Error()
	}
	rsp, err := event.Wait(l.ctx, pid.String())
	if err != nil {
		return nil, err
	}
	if rsp.Ok {
		bytes, err := json.Marshal(rsp.Data)
		if err != nil {
			return nil, err
		}
		return &types.DownstreamCallRsp{
			Data: string(bytes),
		}, nil
	} else {
		errStr, ok := rsp.Data.(string)
		if ok {
			return nil, errors.New(errStr)
		}
		return nil, errors.New("error")
	}
}

func method2Topic(method string) (string, error) {
	method = strings.Replace(method, "/Api/", "", 1)
	topic := ""
	if strings.HasPrefix(method, "PublishTopic") {
		topic = strings.Replace(method, "PublishTopic", "", 1)
		topic = ros.Name2Topic(topic)
	} else if strings.HasPrefix(method, "CallService") {
		topic = strings.Replace(method, "CallService", "", 1)
		topic = ros.Name2Topic(topic)
	} else {
		return "", errors.New(fmt.Sprintf("unknown method %s", method))
	}
	return topic, nil
}
