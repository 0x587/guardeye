package foxglove

import (
	"fmt"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/provider/foxglove/wsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

func New(ip string, port int, topics ...string) provider.IF {
	res := &impl{
		ipPort: fmt.Sprintf("%s:%d", ip, port),
		ws:     wsclient.New(ip, port, topics...),
		out:    make(chan *provider.Msg),
	}
	go res.loop()
	return res
}

type impl struct {
	ipPort string
	ws     *wsclient.Impl
	out    chan *provider.Msg
}

func (i *impl) Get() <-chan *provider.Msg {
	return i.out
}

func (i *impl) loop() {
	go func() {
		err := i.ws.Handle()
		if err != nil {
			logx.Error(err)
		}
	}()
	go func() {
		for msg := range i.ws.GetOutputChan() {
			i.out <- &provider.Msg{
				Message: msg.Data,
				Type:    report.LogType_JSON,
				Provider: report.Provider{
					Type: "RostopicFoxglove",
					Args: []string{i.ipPort, msg.Topic, msg.SchemaName},
				},
			}
		}
	}()
}
