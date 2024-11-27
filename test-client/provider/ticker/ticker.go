package ticker

import (
	"context"
	"time"

	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/provider"
)

func New(ctx context.Context, d time.Duration) provider.IF {
	res := &impl{
		out:    make(chan *provider.Msg),
		d:      d,
		ticker: time.NewTicker(d),
	}
	go res.loop(ctx)
	return res
}

type impl struct {
	d      time.Duration
	ticker *time.Ticker
	out    chan *provider.Msg
}

func (i *impl) loop(ctx context.Context) {
	defer i.ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-i.ticker.C:
			i.out <- &provider.Msg{
				Message:  "ticket in " + t.Format(time.RFC3339Nano),
				Provider: i.getProvider(),
			}
		}
	}
}

func (i *impl) getProvider() reportclient.Provider {
	return reportclient.Provider{
		Type: provider.Ticker,
		Args: []string{i.d.String()},
	}
}

func (i *impl) Get() <-chan *provider.Msg {
	return i.out
}
