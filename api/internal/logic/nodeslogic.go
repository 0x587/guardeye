package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodesLogic {
	return &NodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodesLogic) Nodes(req *types.NodesReq) (resp *types.NodesRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
