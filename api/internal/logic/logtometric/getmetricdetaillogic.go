package logtometric

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/report/report"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMetricDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMetricDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMetricDetailLogic {
	return &GetMetricDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMetricDetailLogic) GetMetricDetail(req *types.GetMetricDetailReq) (resp *types.GetMetricDetailRsp, err error) {
	info, err := l.svcCtx.MetricRedisClient.Info(l.ctx, &report.NodeInfo{ClientId: req.NodeId}, req.Name)
	if err != nil {
		return nil, err
	}
	resp = &types.GetMetricDetailRsp{
		FirstTimestamp: info.FirstTimestamp,
		LastTimestamp:  info.LastTimestamp,
		TotalSamples:   info.TotalSamples,
	}
	return
}
