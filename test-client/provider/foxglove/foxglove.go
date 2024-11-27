package foxglove

import (
	"context"
	"fmt"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/provider/foxglove/wsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

func New(ctx context.Context, ip string, port int, topics ...string) provider.IF {
	res := &impl{
		ipPort: fmt.Sprintf("%s:%d", ip, port),
		ws:     wsclient.New(ip, port, topics...),
		out:    make(chan *provider.Msg),
	}
	go res.loop(ctx)
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

func (i *impl) loop(ctx context.Context) {
	go func() {
		err := i.ws.Handle(ctx)
		if err != nil {
			logx.Error(err)
		}
	}()
	go func() {
		for msg := range i.ws.GetOutputChan() {
			i.out <- &provider.Msg{
				Message: msg.Data,
				Provider: report.Provider{
					Type: provider.Foxglove,
					Args: []string{i.ipPort, msg.Topic, msg.SchemaName},
				},
			}
		}
	}()
}
