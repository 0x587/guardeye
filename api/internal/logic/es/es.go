package es

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
)

const (
	EachFetchCount = 1000
	PitKeepAlive   = "10m"
)

func openPit(ctx context.Context, es *elasticsearch.Client) (string, error) {
	rsp, err := es.OpenPointInTime(
		[]string{"busi-log-*"},
		PitKeepAlive,
		es.OpenPointInTime.WithContext(ctx),
	)
	if err != nil {
		return "", errors.Wrap(err, "Error getting pit response")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(rsp.Body)
	if rsp.IsError() {
		return "", errors.Errorf("Error pit response from server code: %d", rsp.StatusCode)
	}
	var pitRsp struct {
		Id string `json:"id"`
	}
	if err := json.NewDecoder(rsp.Body).Decode(&pitRsp); err != nil {
		return "", errors.Wrap(err, "error in decode pit response")
	}
	return pitRsp.Id, nil
}

func closePit(ctx context.Context, es *elasticsearch.Client, pit string) error {
	var buf bytes.Buffer
	logc.Must(json.NewEncoder(&buf).Encode(struct {
		Id string `json:"id"`
	}{pit}))
	rsp, err := es.ClosePointInTime(
		es.ClosePointInTime.WithContext(ctx),
		es.ClosePointInTime.WithBody(&buf),
	)
	if err != nil {
		return errors.Wrap(err, "Error closing pit response")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(rsp.Body)
	if rsp.IsError() {
		return errors.Errorf("Error closing pit response from server code: %d", rsp.StatusCode)
	}
	return nil
}

func search(ctx context.Context, es *elasticsearch.Client, query M) (*esSearchRes, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, errors.Wrap(err, "Error encoding query")
	}
	rsp, err := es.Search(
		es.Search.WithContext(ctx),
		//es.Search.WithIndex("busi-log-*"), // 使用PIT后不需要再次指定index
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true), // 返回命中总数
	)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting search response")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(rsp.Body)
	// 解析返回结果
	if rsp.IsError() {
		return nil, errors.Errorf("Error search response from server code: %d \nerr: %s", rsp.StatusCode, rsp.String())
	}
	var searchResult esSearchRes
	if err := json.NewDecoder(rsp.Body).Decode(&searchResult); err != nil {
		return nil, errors.Wrap(err, "Error parsing the response body")
	}
	return &searchResult, nil
}
