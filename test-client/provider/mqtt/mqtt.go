package mqtt

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/test-client/provider"
)

func New(ctx context.Context, username, password, broker string) provider.IF {
	res := &impl{
		username: username,
		password: password,
		broker:   broker,
		out:      make(chan *provider.Msg),
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", broker))
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetConnectRetry(true)
	opts.SetClientID("guardeye-agent-" + uuid.New().String())
	opts.SetOnConnectHandler(res.onConnect)
	client := mqtt.NewClient(opts)
	res.mqttClient = client
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logx.Must(token.Error())
	}
	logx.Infof("Connect to %s", broker)

	go func() {
		select {
		case <-ctx.Done():
			logx.Errorf("Disconnect")
			client.Disconnect(1000)
			close(res.out)
		}
	}()
	return res
}

func (i *impl) onConnect(c mqtt.Client) {
	if token := c.Subscribe("#", 0, i.msgHandle); token.Wait() && token.Error() != nil {
		logx.Error(token.Error())
		go func() {
			time.Sleep(time.Second)
			i.onConnect(c)
		}()
	}
	logx.Infof("Subscribe to all topics")
	return
}

type impl struct {
	username   string
	password   string
	broker     string
	out        chan *provider.Msg
	mqttClient mqtt.Client
}

func (i *impl) Get() <-chan *provider.Msg {
	return i.out
}

func (i *impl) msgHandle(_ mqtt.Client, msg mqtt.Message) {
	m := &provider.Msg{
		Message: string(msg.Payload()),
		Provider: report.Provider{
			Type: provider.Mqtt,
			Args: []string{i.broker, i.username, i.password, msg.Topic()},
		},
	}
	select {
	case i.out <- m:
	default:
		logx.Errorf("out is full")
		return
	}
}
