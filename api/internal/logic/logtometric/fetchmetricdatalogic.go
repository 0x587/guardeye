package logtometric

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/report/report"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchMetricDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchMetricDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchMetricDataLogic {
	return &FetchMetricDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchMetricDataLogic) FetchMetricData(req *types.FetchMetricDataReq) (resp *types.FetchMetricDataRsp, err error) {
	o := &metricredis.FetchOption{
		Start: req.From,
		End:   req.To,
	}
	aggMap := map[string]redis.Aggregator{
		"avg":   redis.Avg,
		"sum":   redis.Sum,
		"min":   redis.Min,
		"max":   redis.Max,
		"range": redis.Range,
		"count": redis.Count,
	}
	if agg, ok := aggMap[req.Aggregator]; ok {
		o.Aggregator = agg
		o.WindowSize = time.Duration(req.WindowSecond) * time.Second
	}
	data, err := l.svcCtx.MetricRedisClient.FetchData(
		l.ctx, &report.NodeInfo{ClientId: req.Metric.NodeId}, req.Metric.Name, o)
	if err != nil {
		return nil, err
	}
	res := lo.SliceToMap(data, func(item redis.TSTimestampValue) (int64, float64) {
		return item.Timestamp, item.Value
	})
	//res := lo.Map(data, func(item redis.TSTimestampValue, _ int) types.FetchMetricDataRspData {
	//
	//})
	for i := req.From; i < req.To; i += req.WindowSecond * int(time.Second/time.Millisecond) {
		if _, ok := res[int64(i)]; !ok {
			res[int64(i)] = 0
		}
	}
	resp = &types.FetchMetricDataRsp{
		Data: lo.MapToSlice(res, func(key int64, value float64) types.FetchMetricDataRspData {
			return types.FetchMetricDataRspData{
				Timestamp: key,
				Value:     value,
			}
		}),
	}
	return
}
