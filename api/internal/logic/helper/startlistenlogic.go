package helper

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"unicode"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartListenLogic {
	return &StartListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var listenState = sync.Map{}

func (l *StartListenLogic) StartListen(req *types.ListenReq) error {
	_, ok := listenState.Load(fmt.Sprintf("%s_%s", req.Cid, req.Topic))
	if ok {
		return nil
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	stream, err := l.svcCtx.LinkCli.LinkStreamCall(ctx, &linkclient.LinkCallReq{
		Cid:    req.Cid,
		Method: "SubscribeTopic" + topic2Name(req.Topic),
	})
	if err != nil {
		return err
	}
	listenState.Store(fmt.Sprintf("%s_%s", req.Cid, req.Topic), cancelFunc)
	go func() {
		for {
			_, err := stream.Recv()
			if err != nil {
				listenState.Delete(fmt.Sprintf("%s_%s", req.Cid, req.Topic))
				return
			}
		}
	}()
	return nil
}

func topic2Name(topic string) string {
	var result strings.Builder
	uppercaseNext := false

	for _, r := range topic { // 去掉开头的 `/`
		if r == '/' {
			uppercaseNext = true
		} else {
			if uppercaseNext {
				result.WriteRune(unicode.ToUpper(r))
				uppercaseNext = false
			} else {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}
