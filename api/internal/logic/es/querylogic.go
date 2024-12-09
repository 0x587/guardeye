package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
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
	M  = map[string]interface{}
	MS = []M
)

type esSearchRes struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		Hits []struct {
			Index  string              `json:"_index"`
			Type   string              `json:"_type"`
			Id     string              `json:"_id"`
			Fields map[string][]string `json:"fields"`
			Sort   []int64             `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		TimeAgg struct {
			Buckets []struct {
				KeyAsString time.Time `json:"key_as_string"`
				Key         int64     `json:"key"`
				DocCount    int       `json:"doc_count"`
			} `json:"buckets"`
		} `json:"time_agg"`
	} `json:"aggregations"`
}

func (l *QueryLogic) Query(req *types.EsQueryReq) (resp *types.EsQueryRsp, err error) {
	return nil, nil
}

func (l *QueryLogic) f(req *types.EsQueryReq) (resp *types.EsQueryRsp, err error) {
	es := l.svcCtx.Es
	// 构建搜索查询
	sort := MS{
		{
			"@timestamp": M{
				"order":         "desc",
				"unmapped_type": "boolean",
			},
		},
	}
	filters := MS{
		{
			"range": M{
				"@timestamp": M{
					"gte":    "2024-12-02T07:01:25.181Z",
					"lte":    "2024-12-09T07:01:25.181Z",
					"format": "strict_date_optional_time",
				},
			},
		},
	}
	query := M{
		"size":    500,
		"sort":    sort,
		"_source": false,
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
				"field":            "*",
				"include_unmapped": "true",
			},
			{
				"field":  "@timestamp",
				"format": "strict_date_optional_time",
			},
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	indexName := "busi-log-*" // 替换为你的索引名
	searchRes, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true), // 返回命中总数
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting search response: %s", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(searchRes.Body)
	// 解析返回结果
	if searchRes.IsError() {
		log.Fatalf("Error response from server: %s", searchRes)
	}

	var searchResult esSearchRes
	if err := json.NewDecoder(searchRes.Body).Decode(&searchResult); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// 打印搜索结果
	hits := searchResult.Hits.Hits
	for _, hit := range hits {
		fmt.Printf("%v", hit.Fields)
	}
	return
}
