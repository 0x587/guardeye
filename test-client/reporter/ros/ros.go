package ros

import (
	"encoding/json"
	"errors"
	"os/exec"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ActionSendTopic = "send_topic"
)

func Do(action string, payload map[string]any) (any, error) {
	i := newImpl()
	switch action {
	case ActionSendTopic:
		p := &sendTopicPayload{}
		if err := mapstructure.Decode(payload, p); err != nil {
			return nil, err
		}
		return nil, i.sendTopic(p)
	}
	return nil, errors.New("Unknown action")
}

type impl struct {
	rosBin string
}

func newImpl() *impl {
	return &impl{
		rosBin: "/opt/ros/humble/bin/ros2",
	}
}

type sendTopicPayload struct {
	RosTopic string `mapstructure:"ros_topic"`
	Data     any    `mapstructure:"data"`
}

func (i *impl) sendTopic(p *sendTopicPayload) error {
	yamlData, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	logx.Info("send ", string(yamlData), "to ", p.RosTopic)
	topicType, err := i.getTopicType(p.RosTopic)
	if err != nil {
		return err
	}
	cmd := exec.Command(i.rosBin, "topic", "pub", p.RosTopic, topicType, string(yamlData), "-1")
	//logx.Info("exec ", cmd.String())
	_, err = cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) getTopicType(topic string) (string, error) {
	cmd := exec.Command(i.rosBin, "topic", "type", topic)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s := string(out)
	s = strings.Replace(s, "\n", "", -1)
	return s, nil
}
