package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/0x587/guardeye/report/internal/config"
	"github.com/0x587/guardeye/report/internal/mqs"
	"github.com/0x587/guardeye/report/internal/server"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
)

var configFile = flag.String("f", "etc/report.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := context.Background()
	svcCtx := svc.NewServiceContext(c)
	sg := service.NewServiceGroup()
	defer sg.Stop()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		report.RegisterReportServer(grpcServer, server.NewReportServer(svcCtx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	sg.Add(s)
	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		sg.Add(mq)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	sg.Start()
}
