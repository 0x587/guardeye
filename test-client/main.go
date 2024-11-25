package main

import (
	"context"
	"github.com/0x587/guardeye/test-client/provider/foxglove"
	"os"
	"os/signal"
	"syscall"

	"github.com/0x587/guardeye/test-client/reporter"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//cli := zrpc.MustNewClient(zrpc.RpcClientConf{
	//  Target: "ws.scut.mcurobot.com:56680",
	//})
	cli := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "localhost:8080",
	})
	r := reporter.New(cli,
		//filewatch.New(ctx, "/home/pi/.ros/log/"),
		//ticker.New(ctx, 2500*time.Millisecond),
		//rostopic.New(ctx, "/educar_base_controller/odom"),
		//rostopic.New(ctx, "/camera/camera/image_raw"),
		//rostopic.New(ctx, "/scan"),
		//mqttjson.New("b3351", "scutb3351-mqtt", "ws.scut.mcurobot.com:51883"),
		foxglove.New("10.0.1.109", 8765, "/educar_base_controller/odom"),
	)
	go r.Loop(ctx)
	<-sigChan
	cancel()
}
