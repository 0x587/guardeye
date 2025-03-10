package provider

import (
	"context"
	"errors"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/test-client/config"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/provider/filewatch"
	"github.com/0x587/guardeye/test-client/provider/mqtt"
	"github.com/0x587/guardeye/test-client/provider/rosecho"
	"github.com/0x587/guardeye/test-client/provider/ticker"
)

func New(ctx context.Context, config []config.ProviderConfig) []provider.IF {
	var providers []provider.IF
	for _, p := range config {
		var res provider.IF
		switch p.Type {
		default:
			logx.Must(errors.New("unknown provider type"))
		case provider.Ticker:
			var args struct {
				Interval int
			}
			logx.Must(mapstructure.Decode(p.Args, &args))
			logx.Infof("Provider %s (%d)", p.Type, args.Interval)
			res = ticker.New(ctx, time.Duration(args.Interval)*time.Millisecond)
		case provider.FileWatch:
			var args struct {
				Path string
			}
			logx.Must(mapstructure.Decode(p.Args, &args))
			logx.Infof("Provider %s (%s)", p.Type, args.Path)
			res = filewatch.New(ctx, args.Path)
		case provider.RosEcho:
			var args struct {
				Topic string
			}
			logx.Must(mapstructure.Decode(p.Args, &args))
			logx.Infof("Provider %s (%s)", p.Type, args.Topic)
			res = rosecho.New(ctx, args.Topic)
		case provider.Mqtt:
			var args struct {
				Username string
				Password string
				Broker   string
			}
			logx.Must(mapstructure.Decode(p.Args, &args))
			logx.Infof("Provider %s (%s, %s, %s)", p.Type, args.Username, args.Password, args.Broker)
			res = mqtt.New(ctx, args.Username, args.Password, args.Broker)
		}
		providers = append(providers, res)
	}
	return providers
}
