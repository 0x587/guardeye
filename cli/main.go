package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/0x587/guardeye/link/linkclient"
)

var cid = flag.String("id", "", "agent id")
var outPath = flag.String("o", "", "output path")

func main() {
	flag.Parse()
	logx.MustSetup(logx.LogConf{
		Encoding: "plain",
	})
	client := lo.Must(zrpc.NewClient(zrpc.RpcClientConf{
		Target: "localhost:8080",
	}))
	cli := linkclient.NewLink(client)
	ctx := context.Background()
	genRsp := lo.Must(cli.TypeGen(ctx, &linkclient.TypeGenReq{
		Cid:         *cid,
		TopicRegexp: "",
	}))
	err := os.Mkdir(*outPath, 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	for name, content := range genRsp.GetPb() {
		f := lo.Must(os.Create(fmt.Sprintf("%s/%s.proto", *outPath, name)))
		_ = lo.Must(f.WriteString(content))
	}
}
