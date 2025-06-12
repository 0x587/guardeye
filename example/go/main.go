package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/example/go/pb/pb"
	"github.com/0x587/guardeye/sdk/gosdk"
)

func main() {
	client := gosdk.NewClientConn("linkgrpc.guardeye.shawnsiu.site:5080", "73cf612b-5bff-41f5-bf79-9c2f6698f61d")
	cli := pb.NewApiClient(client)
	ctx := context.Background()
	//for {
	//	rsp, err := cli.CallServiceAddTwoIntsSrv(ctx, &pb.AddReq{
	//		A: &pb.V2{A: 1, B: 2},
	//		B: &pb.V2{A: 3, B: 4},
	//	})
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(rsp)
	//	time.Sleep(time.Second)
	//}
	stream, err := cli.SubscribeTopicV2Publisher(ctx, nil)
	if err != nil {
		logx.Error(err)
	}
	for {
		recv, err := stream.Recv()
		if err != nil {
			logx.Error(err)
			return
		}
		recvJson, _ := json.Marshal(recv)
		fmt.Printf("%s\n", recvJson)
	}
}
