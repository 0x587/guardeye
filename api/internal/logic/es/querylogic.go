package es

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/gql"
	"github.com/0x587/guardeye/common/gql/listener"
	"github.com/0x587/guardeye/common/profile"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type (
	M      = map[string]interface{}
	MS     = []M
	record struct {
		Index  string `json:"_index"`
		Type   string `json:"_type"`
		Id     string `json:"_id"`
		Fields struct {
			Timestamp []time.Time `json:"@timestamp"`
			Message   []string    `json:"log.message"`
		} `json:"fields"`
		Sort []int64 `json:"sort"`
	}
	total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	}
	esSearchRes struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		//Shards   struct {
		//	Total      int `json:"total"`
		//	Successful int `json:"successful"`
		//	Skipped    int `json:"skipped"`
		//	Failed     int `json:"failed"`
		//} `json:"_shards"`
		Hits struct {
			Total total    `json:"total"`
			Hits  []record `json:"hits"`
		} `json:"hits"`
		Aggregations map[string]struct {
			Buckets []struct {
				KeyAsString string `json:"key_as_string"`
				Key         int64  `json:"key"`
				DocCount    int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"aggregations"`
	}
)

func (l *QueryLogic) Query(req *types.EsQueryReq) (resp *types.EsQueryRsp, err error) {
	tree, errs := gql.ParseQuery(req.Query)
	if len(errs) > 0 {
		return &types.EsQueryRsp{
			QueryErrors: errs,
		}, nil
	}
	s := gql.ScheduleQuery(tree)
	filter, err := makeFilter(s)
	if err != nil {
		return nil, err
	}

	records, err, fetchSpend := profile.Measure(func() ([]record, error) {
		recordsCh, err := fetchEs(l.ctx, l.svcCtx.Es, filter)
		res := lo.Map(lo.ChannelToSlice(recordsCh), func(item resInfo, _ int) []record {
			return item.records
		})
		return lo.Flatten(res), err
	})
	if err != nil {
		return nil, err
	}

	type timestamp int
	type row lo.Tuple2[timestamp, []string] // 一行数据
	colNames := lo.Map(s.Result, func(r *listener.ResultEntry, _ int) string { return r.Alias })
	wFunc := lo.Map(s.Result, func(r *listener.ResultEntry, _ int) string { return r.WindowFunc })
	data, err, evalSpend := profile.Measure(func() (map[int][]string, error) {
		// 计算每列的数据
		data := lo.FilterMap(records, func(r record, _ int) (row, bool) {
			ij := gql.NewInjector(r.Fields.Message[0])
			rs := lo.Map(s.Result, func(result *listener.ResultEntry, _ int) string {
				v, err := result.Value.Vf(ij)
				if err != nil {
					if req.TraceError {
						errs = append(errs, err.Error())
					}
					return ""
				}
				return fmt.Sprintf("%v", v)
			})
			flag := false
			for _, i := range rs {
				if i != "" {
					flag = true
					break
				}
			}
			return row{
				A: timestamp(r.Fields.Timestamp[0].In(lo.Must(time.LoadLocation("Asia/Shanghai"))).Unix()),
				B: rs,
			}, flag
		})
		startAt := lo.MinBy(data, func(a row, b row) bool { return a.A < b.A }).A
		groupRes := lo.GroupBy(data, func(r row) int {
			return int(r.A-startAt) / int(s.WindowSize.Seconds())
		})
		fs := lo.Map(wFunc, func(wf string, _ int) func([]string) []string {
			switch wf {
			case "avg":
				return func(vs []string) []string {
					vsAsFloat := lo.Map(vs, func(v string, _ int) float64 {
						f, _ := strconv.ParseFloat(v, 64)
						return f
					})
					avg := lo.Sum(vsAsFloat) / float64(len(vs))
					s := fmt.Sprintf("%f", avg)
					res := make([]string, len(vs))
					res[0] = s
					for i := 1; i < len(vs); i++ {
						res[i] = "inf"
					}
					return res
				}
			}
			return nil
		})
		groupRes = lo.MapValues(groupRes, func(rows []row, k int) []row {
			for colIdx, f := range fs {
				if f == nil {
					continue
				}
				vs := lo.Map(rows, func(r row, _ int) string { return r.B[colIdx] })
				vs = f(vs)
				for rowIdx, r := range rows {
					r.B[colIdx] = vs[rowIdx]
				}
			}
			return lo.Filter(rows, func(r row, _ int) bool {
				for _, s := range r.B {
					if s != "inf" {
						return true
					}
				}
				return false
			})
		})
		res := make(map[int][]string, lo.SumBy(lo.Values(groupRes), func(v []row) int { return len(v) }))
		for _, rows := range groupRes {
			for _, r := range rows {
				res[int(r.A)] = r.B
			}
		}
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	resp = &types.EsQueryRsp{
		Data: lo.MapToSlice(data, func(key int, value []string) types.EsQueryRspData {
			return types.EsQueryRspData{
				Timestamp: key,
				Value:     value,
			}
		}),
		EvalErrors:  errs,
		ColumnNames: colNames,
		Profile: types.EsQueryProfile{
			FetchTime:   int(fetchSpend / time.Millisecond),
			EvalTime:    int(evalSpend / time.Millisecond),
			FetchCount:  len(records),
			ResultCount: len(data),
		},
		//Count:    len(data),
		//RawQuery: string(lo.Must(json.Marshal(query))),
	}
	return
}
