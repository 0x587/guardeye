package es

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/0x587/guardeye/common/gql/listener"
)

func makeFilter(s *listener.Schedule) (MS, error) {
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
	return filter, nil
}

func fetchEs(ctx context.Context, es *elasticsearch.Client, filters MS) (chan resInfo, error) {
	pit, err := openPit(ctx, es)
	if err != nil {
		return nil, errors.Wrap(err, "fail in open pit")
	}

	var searchResult *esSearchRes
	var resCh = make(chan resInfo, 1)
	go func() {

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
		defer func() {
			logc.Infof(ctx, "close")
			if err := closePit(ctx, es, pit); err != nil {
				logc.Error(ctx, err)
			}
			close(resCh)
		}()
		for {
			searchResult, err = search(ctx, es, query)
			if err != nil {
				logc.Error(ctx, errors.Wrap(err, "fail in es search"))
				break
			}
			resCh <- resInfo{
				records: searchResult.Hits.Hits,
				total:   &searchResult.Hits.Total,
			}
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
		return t.Value < EachFetchCount
	case "lt":
		return true
	case "gt":
		return false
	default:
		panic(errors.Errorf("unknow relation %s", t.Relation))
	}
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

type resInfo struct {
	records []record
	total   *total
}
