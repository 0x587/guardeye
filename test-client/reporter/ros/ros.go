package ros

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ActionSendTopic = "send_topic"
)

func Do(action string, payload map[string]any) error {
	switch action {
	case ActionSendTopic:
		p := &sendTopicPayload{}
		if err := mapstructure.Decode(payload, p); err != nil {
			return err
		}
		return sendTopic(p)
	}
	return nil
}

type sendTopicPayload struct {
	RosTopic string `mapstructure:"ros_topic"`
	Data     any    `mapstructure:"data"`
}

func sendTopic(p *sendTopicPayload) error {
	yamlData, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	logx.Info("send ", string(yamlData), "to ", p.RosTopic)
	topicType, err := geTopicType(p.RosTopic)
	if err != nil {
		return err
	}
	cmd := exec.Command("/opt/ros/humble/bin/ros2", "topic", "pub", p.RosTopic, topicType, string(yamlData), "-1")
	logx.Info("exec ", cmd.String())
	return cmd.Run()
}

func geTopicType(topic string) (string, error) {
	cmd := exec.Command("/opt/ros/humble/bin/ros2", "topic", "type", topic)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s := string(out)
	s = strings.Replace(s, "\n", "", -1)
	return s, nil
}
