package ticker

import (
	"context"
	"time"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/provider"
)

func New(ctx context.Context, d time.Duration) provider.IF {
	res := &impl{
		out:    make(chan string),
		d:      d,
		ticker: time.NewTicker(d),
	}
	go func() {
		defer res.ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-res.ticker.C:
				res.out <- t.Format(time.RFC3339Nano)
			}
		}
	}()
	return res
}

type impl struct {
	d      time.Duration
	ticker *time.Ticker
	out    chan string
}

func (i *impl) getProvider() reportclient.Provider {
	return reportclient.Provider{
		Type: ProviderType,
		Args: []string{i.d.String()},
	}
}

const (
	ProviderType = "Ticker"
)

func (i *impl) Get() <-chan *provider.Msg {
	res := make(chan *provider.Msg)
	go func() {
		for msg := range i.out {
			res <- &provider.Msg{
				Message:  msg,
				Type:     report.LogType_TEXT,
				Provider: i.getProvider(),
			}
		}
	}()
	return res
}
