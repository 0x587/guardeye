package main

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/0x587/guardeye/cli/pb/pb"
	"github.com/0x587/guardeye/link/linkclient"
)

type cc struct {
	cli linkclient.Link
}

func (c *cc) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	argsMessage := args.(proto.Message)
	pbBytes, err := proto.Marshal(argsMessage)
	if err != nil {
		return err
	}
	pbBase64 := base64.StdEncoding.EncodeToString(pbBytes)
	callRsp, err := c.cli.LinkCall(ctx, &linkclient.LinkCallReq{
		Cid:    "106bc8ee-6048-4024-8afb-c294ec8fd559",
		Method: method,
		Data:   pbBase64,
	})
	if err != nil {
		return err
	}
	rspBuff, err := base64.StdEncoding.DecodeString(callRsp.Data)
	if err != nil {
		return err
	}
	replyMsg := reply.(proto.Message)
	if err := proto.Unmarshal(rspBuff, replyMsg); err != nil {
		return err
	}
	return nil
}

func (c *cc) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//TODO implement me
	panic("implement me")
}

func newCC() grpc.ClientConnInterface {
	client := lo.Must(zrpc.NewClient(zrpc.RpcClientConf{
		Target: "localhost:8080",
	}))
	cli := linkclient.NewLink(client)
	return &cc{
		cli: cli,
	}
}

func main() {
	logx.MustSetup(logx.LogConf{Encoding: "plain"})
	ctx := context.Background()

	cli := pb.NewApiClient(newCC())

	rsp, err := cli.CallServiceGetPerson(ctx, &pb.GetPersonReq{
		Req: &pb.Person{
			V: &pb.V3{X: 1, Y: 2, Z: 3},
			Pos: &pb.Pos{
				Type: "123",
				V:    &pb.V3{X: 1, Y: 2, Z: 3},
			},
			State: &pb.State{
				Count: 213,
				V:     &pb.V3{X: 1, Y: 2, Z: 3},
			},
		},
	})
	if err != nil {
		logx.Error(err)
		return
	}
	r, _ := json.Marshal(rsp)
	logx.Infof("%s", r)
	chatter, err := cli.PublishTopicChatter(ctx, &pb.String{Data: "123587"})
	if err != nil {
		logx.Error(err)
		return
	}
	r, _ = json.Marshal(chatter)
	logx.Infof("%s", r)
}
