package svc

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/0x587/guardeye/hub/internal/config"
	"github.com/0x587/guardeye/report/reportclient"
)

type ServiceContext struct {
	Config   config.Config
	Upstream reportclient.Report
}

func NewServiceContext(c config.Config) *ServiceContext {
	opts := zrpc.RpcClientConf{}
	logx.Must(conf.FillDefault(&opts))
	opts.Endpoints = c.Upstream.Endpoints
	cli := zrpc.MustNewClient(opts)
	return &ServiceContext{
		Config:   c,
		Upstream: reportclient.NewReport(cli),
	}
}
