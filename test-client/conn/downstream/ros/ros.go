package ros

import (
	"github.com/0x587/guardeye/link/linkclient"
)

const (
	ActionSendTopic   = "send_topic"
	ActionCallService = "call_service"
)

type IF interface {
	Exec(payload *linkclient.LinkCommandPayloadRosExec) ([]byte, error)
	Type(payload *linkclient.LinkCommandPayloadRosType) (string, string, error)
}
