package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0x587/guardeye/test-client/provider/filewatch"
	"github.com/0x587/guardeye/test-client/provider/rostopic"
	"github.com/0x587/guardeye/test-client/provider/ticker"
	"github.com/0x587/guardeye/test-client/reporter"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fileProvider := filewatch.New(ctx, "./test")
	tickerProvider := ticker.New(ctx, 2500*time.Millisecond)
	rosTopicProvider := rostopic.New(ctx, "/educar_base_controller/odom")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//cli := zrpc.MustNewClient(zrpc.RpcClientConf{
	//	Target: "ws.scut.mcurobot.com:50080",
	//})
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "10.0.1.103:8080",
	})
	r := reporter.New(cli, fileProvider, tickerProvider, rosTopicProvider)
	go r.Loop(ctx)
	<-sigChan
	cancel()
}
