package helper

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
)

type StopListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStopListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StopListenLogic {
	return &StopListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StopListenLogic) StopListen(req *types.ListenReq) error {
	v, ok := listenState.Load(fmt.Sprintf("%s_%s", req.Cid, req.Topic))
	if !ok {
		return nil
	}
	cancelFunc := v.(context.CancelFunc)
	cancelFunc()
	listenState.Delete(fmt.Sprintf("%s_%s", req.Cid, req.Topic))
	return nil
}
