package logtometric

//import (
//	"github.com/0x587/guardeye/api/internal/types"
//	"github.com/0x587/guardeye/report/report"
//	"github.com/samber/lo"
//)

//func MetricQueryConvert(mq *report.MetricQuery) types.MetricQuery {
//	return types.MetricQuery{
//		Type:      mq.GetType().String(),
//		Name:      mq.GetName(),
//		ParsePath: mq.GetParsePath(),
//		Filters: lo.Map(mq.GetFilters(), func(mf *report.MetricFilter, _ int) types.MetricFilter {
//			return MetricFilterConvert(mf)
//		}),
//	}
//}
//
//func MetricFilterConvert(mf *report.MetricFilter) types.MetricFilter {
//	return types.MetricFilter{
//		Metric: MetricQueryConvert(mf.GetMetric()),
//		Op:     mf.GetOp().String(),
//		Value:  mf.GetValue(),
//	}
//}
