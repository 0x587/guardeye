package main

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/ttt/foxglove/foxgloveclient"
)

func main() {
	ctx := context.Background()
	fgCli := foxgloveclient.New("127.0.0.1", 8765, "/chatter")
	go func() {
		err := fgCli.Run(ctx)
		if err != nil {
			logx.Error(err)
		}
	}()
	for m := range fgCli.GetOutputChan() {
		fmt.Println(m)
	}
}
