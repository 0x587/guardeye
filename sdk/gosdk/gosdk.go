package gosdk

import (
	"context"
	"encoding/base64"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/0x587/guardeye/link/linkclient"
)

type impl struct {
	cli linkclient.Link
}

func NewClientConn(target string) grpc.ClientConnInterface {
	client := lo.Must(zrpc.NewClient(zrpc.RpcClientConf{
		Target: target,
	}))
	cli := linkclient.NewLink(client)
	return &impl{
		cli: cli,
	}
}

func (i *impl) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	argsMessage := args.(proto.Message)
	pbBytes, err := proto.Marshal(argsMessage)
	if err != nil {
		return err
	}
	pbBase64 := base64.StdEncoding.EncodeToString(pbBytes)
	callRsp, err := i.cli.LinkCall(ctx, &linkclient.LinkCallReq{
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

func (i *impl) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//TODO implement me
	panic("implement me")
}
