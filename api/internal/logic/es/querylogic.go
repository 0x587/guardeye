package es

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/logic/es/listener"
	"github.com/0x587/guardeye/api/internal/logic/es/parser"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
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

func parseQuery(query string) (parser.IParseContext, []string) {
	input := antlr.NewInputStream(query)
	lexer := parser.NewGuardQueryLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGuardQueryParser(stream)
	el := listener.NewErrListener()
	p.AddErrorListener(antlr.NewProxyErrorListener([]antlr.ErrorListener{
		el,
		antlr.NewDiagnosticErrorListener(false),
	}))
	return p.Parse(), el.GetErrs()
}

func scheduleQuery(tree parser.IParseContext) *listener.Schedule {
	s := listener.NewScheduler()
	antlr.NewParseTreeWalker().Walk(s.Listener, tree)
	return s.GetSchedule()
}

func (l *QueryLogic) Query(req *types.EsQueryReq) (resp *types.EsQueryRsp, err error) {
	tree, errs := parseQuery(req.Query)
	if len(errs) > 0 {
		return &types.EsQueryRsp{
			QueryErrors: errs,
		}, nil
	}
	s := scheduleQuery(tree)
	var sourceQuery MS
	for _, source := range s.SourceWhere {
		var qs MS = nil
		if !source.Node.Any {
			qs = append(qs, makeHas("nodeInfo.clientId", source.Node.Nid))
		}
		for _, p := range source.Providers {
			if p.Any {
				continue
			}
			if p.PType != "" {
				qs = append(qs, makeHas("log.provider.type", p.PType))
			}
		}
		for _, key := range source.NeedKeys {
			qs = append(qs, makeHas("log.message", key))
		}
		if qs != nil {
			sourceQuery = append(sourceQuery, makeAnd(qs...))
		}
	}
	from, err := s.TimeWhere.From.Get()
	if err != nil {
		return nil, err
	}
	to, err := s.TimeWhere.To.Get()
	if err != nil {
		return nil, err
	}
	filter := MS{
		{
			"range": M{
				"@timestamp": M{
					"gte":    from.Format("2006-01-02T15:04:05Z07"),
					"lte":    to.Format("2006-01-02T15:04:05Z07"),
					"format": "strict_date_optional_time",
				},
			},
		},
		makeOr(sourceQuery...),
	}
	{
		//	TODO: Just for debug
		b := lo.Must(json.MarshalIndent(filter, "", "  "))
		fmt.Println(string(b))
	}

	records, err, fetchSpend := profile.Measure(func() ([]record, error) {
		recordsCh, err := l.fetchEs(filter)
		return lo.Flatten(lo.ChannelToSlice(recordsCh)), err
	})

	if err != nil {
		return nil, err
	}

	data, _, evalSpend := profile.Measure(func() (map[int][]string, error) {
		data := lo.FilterMap(records, func(r record, _ int) (lo.Tuple2[int, []string], bool) {
			ij := &injector{msg: r.Fields.Message[0]}
			row := lo.Map(s.Result, func(result *listener.ResultEntry, _ int) string {
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
			for _, i := range row {
				if i != "" {
					flag = true
					break
				}
			}
			return lo.Tuple2[int, []string]{A: int(r.Fields.Timestamp[0].In(lo.Must(time.LoadLocation("Asia/Shanghai"))).Unix()), B: row}, flag
		})
		return lo.SliceToMap(data, func(i lo.Tuple2[int, []string]) (int, []string) { return i.A, i.B }), nil
	})
	resp = &types.EsQueryRsp{
		Data:        data,
		EvalErrors:  errs,
		ColumnNames: lo.Map(s.Result, func(r *listener.ResultEntry, _ int) string { return r.Alias }),
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

func (l *QueryLogic) fetchEs(filters MS) (chan []record, error) {
	es := l.svcCtx.Es

	pit, err := openPit(l.ctx, es)
	if err != nil {
		return nil, err
	}

	// 构建搜索查询
	sort := MS{
		{
			"@timestamp": M{
				"order":         "desc",
				"unmapped_type": "boolean",
			},
		},
	}
	query := M{
		"size": EachFetchCount,
		"sort": sort,
		"query": M{
			"bool": M{
				"filter": filters,
			},
		},
		"aggs": M{
			"time_agg": M{
				"date_histogram": M{
					"field":          "@timestamp",
					"fixed_interval": "3h",
					"time_zone":      "Asia/Shanghai",
					"min_doc_count":  1,
				},
			},
		},
		"fields": MS{
			{
				"field": "log.message",
			},
			{
				"field":  "@timestamp",
				"format": "strict_date_optional_time",
			},
		},
		"pit": M{
			"id":         pit,
			"keep_alive": PitKeepAlive,
		},
	}
	var searchResult *esSearchRes
	var resCh = make(chan []record)

	go func() {
		defer func() {
			logc.Infof(l.ctx, "close")
			if err := closePit(l.ctx, es, pit); err != nil {
				logc.Error(l.ctx, err)
			}
			close(resCh)
		}()
		for {
			searchResult, err = search(l.ctx, es, query)
			if err != nil {
				logc.Error(l.ctx, errors.Wrap(err, "fail in es search"))
				break
			}
			resCh <- searchResult.Hits.Hits
			if isAll(searchResult.Hits.Total) {
				break
			} else {
				if len(searchResult.Hits.Hits) == 0 {
					break
				}
				lastRecord := searchResult.Hits.Hits[len(searchResult.Hits.Hits)-1]
				query["search_after"] = lastRecord.Sort
			}
		}
	}()
	return resCh, nil
}

func isAll(t total) bool {
	switch t.Relation {
	case "eq":
		return t.Value <= EachFetchCount
	case "lt":
		return true
	case "gt":
		return false
	default:
		panic(errors.Errorf("unknow relation %s", t.Relation))
	}
}

type injector struct {
	msg string
}

func (i *injector) GetMsg() string {
	return i.msg
}

func makeHas(key, value string) M {
	return M{
		"bool": M{
			"should": MS{
				{
					"match": M{
						key: value,
					},
				},
			},
			"minimum_should_match": 1,
		},
	}
}

//func makeLike(key, value string) M {
//	return M{
//		"bool": M{
//			"should": MS{
//				{
//					"query_string": M{
//						"fields": []string{key},
//						"query":  fmt.Sprintf("* %s", value),
//					},
//				},
//			},
//			"minimum_should_match": 1,
//		},
//	}
//}

func makeAnd(items ...M) M {
	if len(items) == 1 {
		return items[0]
	}
	return M{
		"bool": M{
			"filter": items,
		},
	}
}

func makeOr(items ...M) M {
	if len(items) == 1 {
		return items[0]
	}
	return M{
		"bool": M{
			"should": items,
		},
	}
}
