package main

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/example/go/pb/pb"
	"github.com/0x587/guardeye/sdk/gosdk"
)

func main() {
	client := gosdk.NewClientConn("linkgrpc.guardeye.shawnsiu.site:5080", "59af85f5-244b-4f77-93a9-2d11e86af19b")
	cli := pb.NewApiClient(client)
	ctx := context.Background()
	state, err := cli.CallServiceArebotTransportRobotBridgeGetState(ctx, &pb.GetStateReq{})
	if err != nil {
		return
	}
	logx.Info(state)
	//stream, err := cli.SubscribeTopicArebotTransportRobotBridgeRobotState(ctx, nil)
	//if err != nil {
	//	logx.Error(err)
	//}
	//for {
	//	recv, err := stream.Recv()
	//	if err != nil {
	//		logx.Error(err)
	//		return
	//	}
	//	recvJson, _ := json.Marshal(recv)
	//	fmt.Printf("%s\n", recvJson)
	//	time.Sleep(time.Second)
	//}
}
