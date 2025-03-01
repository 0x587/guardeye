package main

import (
	"context"
	"errors"
	"flag"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/test-client/provider"
	"github.com/0x587/guardeye/test-client/provider/filewatch"
	"github.com/0x587/guardeye/test-client/provider/mqtt"
	"github.com/0x587/guardeye/test-client/provider/rosecho"
	"github.com/0x587/guardeye/test-client/provider/ticker"
)

type configInFile struct {
	Server struct {
		Host string
		Port int
	}
	Provider []struct {
		Type string
		Args map[string]interface{}
	}
}

type Config struct {
	Server struct {
		Host string
		Port int
	}
	Providers []provider.IF
}

var configFile = flag.String("f", "/etc/guardeye-agent/config.yaml", "the config file")

func LoadConfig(ctx context.Context) Config {
	var fc configInFile
	conf.MustLoad(*configFile, &fc)
	var c Config
	c.Server = fc.Server
	for _, p := range fc.Provider {
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
		c.Providers = append(c.Providers, res)
	}
	return c
}
