package logic

import (
	"context"
	"time"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/model"
	"github.com/google/uuid"
	"github.com/samber/lo"

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
	rawlogs, err := l.svcCtx.RawLogDBClient.ListLastSeen(l.ctx)
	if err != nil {
		return nil, err
	}
	lastSeen := lo.SliceToMap(rawlogs, func(item *model.Rawlog) (uuid.UUID, time.Time) {
		return item.ClientId, item.CreatedAt
	})
	resp = &types.NodesRsp{
		Nodes: lo.Map(nodes, func(n *model.Node, index int) types.NodeInfo {
			return types.NodeInfo{
				Id:         n.ClientId.String(),
				Alias:      n.Alias.String,
				Macs:       n.Macs,
				Ips:        n.Macs,
				LastSeenAt: lastSeen[n.ClientId].String(),
			}
		}),
	}
	return
}
