package implfg

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/common/foxgloveclient"
	"github.com/0x587/guardeye/link/link"
	"github.com/0x587/guardeye/link/linkclient"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
)

func New(ip string, port int, patterns []string, callback func(string, []byte)) ros.IF {
	return &impl{
		cli:      foxgloveclient.New(ip, port, patterns).Run(context.Background()),
		callback: callback,
	}
}

type impl struct {
	cli      foxgloveclient.IF
	callback func(string, []byte)
}

func (i *impl) Exec(payload *linkclient.LinkCommandPayloadRosExec) ([]byte, error) {
	switch payload.GetAction() {
	case ros.ActionSendTopic:
		return i.sendTopic(payload)
	case ros.ActionSubscribeTopic:
		return i.subscribeTopic(payload)
	case ros.ActionCallService:
		return i.callService(payload)
	}
	return nil, errors.New("unknown action")
}

func (i *impl) sendTopic(p *linkclient.LinkCommandPayloadRosExec) ([]byte, error) {
	cdrData, err := base64.StdEncoding.DecodeString(p.Data)
	if err != nil {
		return nil, err
	}
	err = i.cli.Publish(p.RosTopic, cdrData)
	if err != nil {
		return nil, err
	}
	return cdrData, nil
}

func (i *impl) subscribeTopic(p *linkclient.LinkCommandPayloadRosExec) ([]byte, error) {
	logx.Errorf("订阅 %v", p)
	err := i.cli.Subscribe(p.GetRosTopic(), func(transData []byte) {
		id := p.GetData() // 请求的轮询ID
		i.callback(id, transData)
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (i *impl) callService(p *linkclient.LinkCommandPayloadRosExec) ([]byte, error) {
	cdrData, err := base64.StdEncoding.DecodeString(p.Data)
	if err != nil {
		return nil, err
	}
	rspCdr, err := i.cli.Call(p.RosTopic, cdrData)
	if err != nil {
		return nil, err
	}
	return rspCdr, nil
}

func (i *impl) Type(payload *linkclient.LinkCommandPayloadRosType) (*linkclient.LinkTypeGenResult, error) {
	if payload.GetType() == link.RosTypeGenType_MESSAGE {
		messageSchema, err := i.cli.GetMessageType(payload.GetRosTopic())
		if err != nil {
			return nil, err
		}
		return &linkclient.LinkTypeGenResult{
			Name: strings.ReplaceAll(messageSchema.Name, "/msg", ""),
			Req:  messageSchema.Schema,
		}, nil
	}
	if payload.GetType() == link.RosTypeGenType_SERVICE {
		serviceSchema, err := i.cli.GetServiceType(payload.GetRosTopic())
		if err != nil {
			return nil, err
		}
		return &linkclient.LinkTypeGenResult{
			Name: strings.ReplaceAll(serviceSchema.Name, "/srv", ""),
			Req:  serviceSchema.RequestSchema,
			Rsp:  serviceSchema.ResponseSchema,
		}, nil
	}
	return nil, errors.New("unknown ros type gen type")
}

func (i *impl) List(payload *linkclient.LinkCommandPayloadRosList) (*linkclient.TypeListRsp, error) {
	return i.cli.List(), nil
}
