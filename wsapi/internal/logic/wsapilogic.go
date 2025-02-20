package logic

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/wsapi/internal/svc"
)

type WsapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWsapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WsapiLogic {
	return &WsapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WsapiLogic) Wsapi(w http.ResponseWriter, r *http.Request) error {
	l.svcCtx.Ws.ServeWs(w, r)
	return nil
}
