package logic

import (
	"context"
	"net/http"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type WsListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsListenLogic {
	return &WsListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsListenLogic) WsListen(w http.ResponseWriter, r *http.Request) error {
	l.svcCtx.ListenWs.ServeWs(w, r)
	return nil
}
