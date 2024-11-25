package logtometric

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/common/tokv"
	"github.com/0x587/guardeye/common/utils"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/samber/lo"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/collection"
	"google.golang.org/protobuf/encoding/protojson"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	return &impl{
		ctx:        ctx,
		svcCtx:     svcCtx,
		queryCache: lo.Must(collection.NewCache(2500*time.Millisecond, collection.WithName("LogToQueryCache"))),
	}
}

type impl struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	queryCache *collection.Cache
}

func (i *impl) Consume(ctx context.Context, _, val string) error {
	d := &report.MQLog{}
	if err := json.Unmarshal([]byte(val), d); err != nil {
		return err
	}
	metricQueries, err := i.getQuery(ctx, d)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil
		}
		return err
	}
	for _, metricQuery := range metricQueries {
		if metricQuery.Type == report.MetricType_String {
			return errors.New("metric type not supported")
		}
		needParseMetrics := append(lo.Map(metricQuery.Filters,
			func(item *report.MetricFilter, _ int) *report.MetricQuery {
				return item.Metric
			}), metricQuery)
		metricsKv, err := parseMetricFromLog(d.GetLog(), needParseMetrics)
		if err != nil {
			return nil
		}
		if !doFilter(metricQuery.Filters, metricsKv) {
			return nil
		}

		v := metricsKv[metricQuery.ParsePath]
		vFloat, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		if err = i.svcCtx.MetricRedisClient.AddMetric(ctx, d.GetNodeInfo(), metricQuery.Name, vFloat); err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) getQuery(ctx context.Context, log *report.MQLog) ([]*report.MetricQuery, error) {
	cid, p := log.GetNodeInfo().GetClientId(), utils.ProviderToStr(log.GetLog().GetProvider())
	key := fmt.Sprintf("%s-%s", cid, p)
	res, err := i.queryCache.Take(key, func() (any, error) {
		ltms, err := i.svcCtx.LogToMetricDBClient.ListForLog(ctx, cid, p)
		if err != nil {
			return nil, err
		}
		mqs := lo.FilterMap(ltms, func(item *model.LogToMetricQuery, index int) (*report.MetricQuery, bool) {
			mq := &report.MetricQuery{}
			if err := protojson.Unmarshal([]byte(item.Query), mq); err != nil {
				return nil, false
			}
			return mq, true
		})
		return mqs, nil
	})
	if err != nil {
		return nil, err
	}
	return res.([]*report.MetricQuery), nil
}

// parseMetricFromLog return {parsePath: value}, any metric parse error will be returned
func parseMetricFromLog(log *report.Log, ms []*report.MetricQuery) (map[string]string, error) {
	var kv tokv.KV
	switch log.GetType() {
	default:
		return nil, errors.New("unknown log type")
	case report.LogType_TEXT:
		return nil, errors.New("text log type not supported")
	case report.LogType_YAML:
		kv = tokv.YamlToKv(log.GetMessage())
	case report.LogType_JSON:
		kv = tokv.JsonToKv(log.GetMessage())
	}
	res := make(map[string]string)
	for _, m := range ms {
		v, ok := kv[m.ParsePath]
		if !ok {
			return nil, errors.New(fmt.Sprintf("parse path not found: %s", m.ParsePath))
		}
		res[m.ParsePath] = v
	}
	return res, nil
}

func doFilter(filters []*report.MetricFilter, kv map[string]string) bool {
	for _, filter := range filters {
		leftValue, ok := kv[filter.Metric.ParsePath]
		if !ok {
			return false
		}
		rightValue := filter.Value

		switch filter.Op {
		default:
			return false
		case report.MetricFilterOp_EQ:
			return leftValue == rightValue
		case report.MetricFilterOp_NE:
			return leftValue != rightValue
		}
	}
	return true
}
