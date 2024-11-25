package logtometric

import (
	"context"
	"errors"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/report"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetQueriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetQueriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetQueriesLogic {
	return &GetQueriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetQueriesLogic) GetQueries(req *types.GetQueriesReq) (resp *types.GetQueriesRsp, err error) {
	resp = &types.GetQueriesRsp{}
	var queries []*model.LogToMetricQuery
	if req.NodeId == "" {
		queries, err = l.svcCtx.LogToMetricDBClient.List(l.ctx)
	} else {
		queries, err = l.svcCtx.LogToMetricDBClient.ListForNode(l.ctx, req.NodeId)
	}
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	type NodeProvider lo.Tuple2[string, string]
	type NodeProviderQuery lo.Tuple2[NodeProvider, types.MetricQuery]
	qs := lo.FilterMap(queries, func(q *model.LogToMetricQuery, _ int) (res NodeProviderQuery, flg bool) {
		mq := &report.MetricQuery{}
		err = protojson.Unmarshal([]byte(q.Query), mq)
		if err != nil {
			logc.Errorf(l.ctx, "query: %s, err: %v", q.Query, err)
			return
		}
		flg = true
		res.A = NodeProvider{A: q.ClientId.String(), B: q.Provider}
		res.B = MetricQueryConvert(mq)
		return
	})
	qsGroup := lo.GroupBy(qs, func(item NodeProviderQuery) NodeProvider { return item.A })
	resp.Queries = lo.MapToSlice(qsGroup,
		func(key NodeProvider, value []NodeProviderQuery) types.MetricQueryGroup {
			return types.MetricQueryGroup{
				NodeId:   key.A,
				Provider: key.B,
				Queries:  lo.Map(value, func(item NodeProviderQuery, _ int) types.MetricQuery { return item.B }),
			}
		})
	return
}
