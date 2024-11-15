package logic

import (
	"context"
	"net/http"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/ws"
	"github.com/zeromicro/go-zero/core/logx"
)

type NodeWsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodeWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodeWsLogic {
	return &NodeWsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodeWsLogic) NodeWs(req *types.NodeWsReq, w http.ResponseWriter, r *http.Request) error {
	bws := l.svcCtx.BoardCaseWs[req.Id]
	if bws == nil {
		bws = ws.New[string]()
		l.svcCtx.BoardCaseWs[req.Id] = bws
	}
	bws.ServeWs(w, r)
	return nil
}
