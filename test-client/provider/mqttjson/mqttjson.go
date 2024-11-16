package mqttjson

import (
	"fmt"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/test-client/provider"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/logc"
)

func New(username, password, broker string) provider.IF {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", broker))
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logc.Must(token.Error())
	}
	res := &impl{
		username:   username,
		password:   password,
		broker:     broker,
		out:        make(chan mqtt.Message),
		mqttClient: client,
	}
	token := client.Subscribe("#", 0, res.msgHandle)
	token.Wait()
	return res
}

type impl struct {
	username   string
	password   string
	broker     string
	out        chan mqtt.Message
	mqttClient mqtt.Client
}

func (i *impl) Get() <-chan *provider.Msg {
	res := make(chan *provider.Msg)
	go func() {
		for msg := range i.out {
			res <- &provider.Msg{
				Message: string(msg.Payload()),
				Type:    report.LogType_JSON,
				Provider: report.Provider{
					Type: "MqttJson",
					Args: []string{i.broker, i.username, i.password, msg.Topic()},
				},
			}
		}
	}()
	return res
}

func (i *impl) msgHandle(_ mqtt.Client, msg mqtt.Message) {
	i.out <- msg
}
