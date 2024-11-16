package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/model"
	"github.com/samber/lo/parallel"

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
	nodes, err := l.svcCtx.NodeDBClient.ListGroupByClientID(l.ctx)
	if err != nil {
		return nil, err
	}
	resp = &types.NodesRsp{
		Nodes: parallel.Map(nodes, func(n *model.Node, index int) types.NodeInfo {
			res := types.NodeInfo{
				Id:    n.ClientId.String(),
				Alias: n.Alias.String,
				Macs:  n.Macs,
				Ips:   n.Macs,
			}
			seen, err := l.svcCtx.RawLogDBClient.GetLastSeen(l.ctx, n.ClientId)
			if err == nil {
				res.LastSeenAt = seen.CreatedAt.String()
			}
			return res
		}),
	}
	return
}
