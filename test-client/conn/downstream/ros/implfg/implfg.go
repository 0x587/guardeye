package implfg

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/0x587/guardeye/common/foxgloveclient"
	"github.com/0x587/guardeye/link/link"
	"github.com/0x587/guardeye/link/linkclient"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
)

func New(ip string, port int) ros.IF {
	return &impl{
		cli: foxgloveclient.New(ip, port).Run(context.Background()),
	}
}

type impl struct {
	cli foxgloveclient.IF
}

func (i *impl) Exec(payload *linkclient.LinkCommandPayloadRosExec) ([]byte, error) {
	switch payload.GetAction() {
	case ros.ActionSendTopic:
		return i.sendTopic(payload)
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

func (i *impl) Type(payload *linkclient.LinkCommandPayloadRosType) (string, string, error) {
	if payload.GetType() == link.RosTypeGenType_MESSAGE {
		messageSchema, err := i.cli.GetMessageType(payload.GetRosTopic())
		if err != nil {
			return "", "", err
		}
		return messageSchema.Schema, "", nil
	}
	if payload.GetType() == link.RosTypeGenType_SERVICE {
		serviceSchema, err := i.cli.GetServiceType(payload.GetRosTopic())
		if err != nil {
			return "", "", err
		}
		return serviceSchema.RequestSchema, serviceSchema.ResponseSchema, nil
	}
	return "", "", errors.New("unknown ros type gen type")
}
