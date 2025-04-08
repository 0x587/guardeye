package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/handler"
	"github.com/0x587/guardeye/api/internal/mqs"
	"github.com/0x587/guardeye/api/internal/svc"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.MustSetup(logx.LogConf{Encoding: "plain"})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	sg := service.NewServiceGroup()
	defer sg.Stop()

	ctx := context.Background()
	svcCtx := svc.NewServiceContext(c)

	apiServer := rest.MustNewServer(c.RestConf, rest.WithCors("*"))
	sg.Add(apiServer)
	handler.RegisterHandlers(apiServer, svcCtx)
	apiServer.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodGet,
				Path:   "/ws/test-page",
				Handler: func(w http.ResponseWriter, r *http.Request) {
					http.ServeFile(w, r, "home.html")
				},
			},
		},
	)

	for _, mq := range mqs.Consumers(c, ctx, svcCtx) {
		break
		sg.Add(mq)
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	sg.Start()
}
