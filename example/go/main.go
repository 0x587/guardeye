package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/example/go/pb/pb"
	"github.com/0x587/guardeye/sdk/gosdk"
)

func main() {
	client := gosdk.NewClientConn("linkgrpc.guardeye.shawnsiu.site:5080", "73cf612b-5bff-41f5-bf79-9c2f6698f61d")
	cli := pb.NewApiClient(client)
	for {
		ctx := context.Background()
		n := time.Now()
		person, err := cli.CallServiceAddTwoIntsSrv(ctx, &pb.AddReq{
			A: &pb.V2{A: int64(n.Hour()), B: int64(n.Minute())},
			B: &pb.V2{A: int64(n.Second()), B: int64(n.Nanosecond())},
		})
		if err != nil {
			logx.Error(err)
		}
		personJson, _ := json.Marshal(person)
		fmt.Printf("%s\n", personJson)
		time.Sleep(time.Second)
	}

}
