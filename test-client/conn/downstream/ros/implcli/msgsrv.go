package implcli

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/0x587/guardeye/link/linkclient"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rossrv"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rossrv/walk"
)

func (i *impl) sendTopic(p *linkclient.LinkCommandPayloadRosExec) (any, error) {
	jsonData, err := json.Marshal(p.Data)
	if err != nil {
		return nil, err
	}
	topicType, err := i.getTopicType(p.RosTopic)
	if err != nil {
		return nil, err
	}
	_, err = i.rosExec(fmt.Sprintf("ros2 topic pub %s %s '%s' -1", p.RosTopic, topicType, string(jsonData)))
	if err != nil {
		return nil, err
	}
	return p.Data, nil
}

func (i *impl) callService(p *linkclient.LinkCommandPayloadRosExec) (any, error) {
	jsonData, err := json.Marshal(p.Data)
	if err != nil {
		return nil, err
	}
	topicType, err := i.getServiceType(p.RosTopic)
	if err != nil {
		return nil, err
	}
	data := string(jsonData)
	cmd := fmt.Sprintf("ros2 service call %s %s '%s'", p.RosTopic, topicType, data)
	res, err := i.rosExec(cmd)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(res, "response:") {
		return nil, errors.New("srv response format wrong")
	}
	parts := strings.Split(res, "response:")
	if len(parts) != 2 {
		return nil, errors.New("srv response format wrong")
	}
	tree, err := rossrv.Parse(parts[1])
	if err != nil {
		return nil, err
	}
	return walk.Walk(tree), nil
}

var topicTypeCache sync.Map

func (i *impl) getTopicType(topic string) (string, error) {
	c, ok := topicTypeCache.Load(topic)
	if ok {
		return c.(string), nil
	}
	s, err := i.rosExec(fmt.Sprintf("ros2 topic type %s ", topic))
	if err != nil {
		return "", err
	}
	s = strings.Replace(s, "\n", "", -1)
	topicTypeCache.Store(topic, s)
	return s, nil
}

var serviceTypeCache sync.Map

func (i *impl) getServiceType(topic string) (string, error) {
	c, ok := serviceTypeCache.Load(topic)
	if ok {
		return c.(string), nil
	}
	s, err := i.rosExec(fmt.Sprintf("ros2 service type %s ", topic))
	if err != nil {
		return "", err
	}
	s = strings.Replace(s, "\n", "", -1)
	serviceTypeCache.Store(topic, s)
	return s, nil
}
