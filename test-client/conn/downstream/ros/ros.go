package ros

import (
	"github.com/0x587/guardeye/link/linkclient"
)

const (
	ActionSendTopic      = "send_topic"
	ActionSubscribeTopic = "subscribe_topic"
	ActionCallService    = "call_service"
)

type IF interface {
	Exec(payload *linkclient.LinkCommandPayloadRosExec) ([]byte, error)
	List(payload *linkclient.LinkCommandPayloadRosList) (*linkclient.TypeListRsp, error)
	Type(payload *linkclient.LinkCommandPayloadRosType) (*linkclient.LinkTypeGenResult, error)
}
