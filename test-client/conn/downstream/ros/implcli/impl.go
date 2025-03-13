package implcli

import (
	"errors"

	"github.com/0x587/guardeye/link/linkclient"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
)

func (i *impl) Do(action string, payload *linkclient.LinkCommandPayloadRosExec) (any, error) {
	switch action {
	case ros.ActionSendTopic:
		return i.sendTopic(payload)
	case ros.ActionCallService:
		return i.callService(payload)
	}
	return nil, errors.New("unknown action")
}

type impl struct{}

//func New() ros.IF {
//	return &impl{}
//}
