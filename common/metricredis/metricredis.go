package metricredis

import (
	"context"
	"fmt"
	"time"

	"github.com/0x587/guardeye/report/report"
	"github.com/redis/go-redis/v9"
)

type TimeSeriesInfo struct {
	TotalSamples   int64
	FirstTimestamp int64
	LastTimestamp  int64
}

type FetchOption struct {
	Start      int
	End        int
	Aggregator redis.Aggregator
	WindowSize time.Duration
}

type IF interface {
	AddMetric(ctx context.Context, nodeInfo *report.NodeInfo, name string, value float64) error
	Info(ctx context.Context, nodeInfo *report.NodeInfo, name string) (*TimeSeriesInfo, error)
	FetchData(ctx context.Context, nodeInfo *report.NodeInfo, name string, o *FetchOption) ([]redis.TSTimestampValue, error)
}

func New(cli *redis.Client) IF {
	return &impl{
		cli: cli,
	}
}

type impl struct {
	cli *redis.Client
}

func (i *impl) FetchData(ctx context.Context, nodeInfo *report.NodeInfo, name string, o *FetchOption) ([]redis.TSTimestampValue, error) {
	ro := &redis.TSRangeOptions{
		Aggregator:     o.Aggregator,
		BucketDuration: int(o.WindowSize / time.Millisecond),
	}
	if ro.Aggregator != 0 {
		ro.Align = "-"
	}
	cmd := i.cli.TSRangeWithArgs(ctx, i.key(nodeInfo, name), o.Start, o.End, ro)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return cmd.Val(), nil
}

func (i *impl) Info(ctx context.Context, nodeInfo *report.NodeInfo, name string) (*TimeSeriesInfo, error) {
	cmd := i.cli.TSInfo(ctx, i.key(nodeInfo, name))
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}
	return &TimeSeriesInfo{
		TotalSamples:   cmd.Val()["totalSamples"].(int64),
		FirstTimestamp: cmd.Val()["firstTimestamp"].(int64),
		LastTimestamp:  cmd.Val()["lastTimestamp"].(int64),
	}, nil
}

func (i *impl) AddMetric(ctx context.Context, nodeInfo *report.NodeInfo, name string, value float64) error {
	cmd := i.cli.TSAdd(ctx, i.key(nodeInfo, name), time.Now().UnixMilli(), value)
	return cmd.Err()
}

func (i *impl) key(nodeInfo *report.NodeInfo, name string) string {
	return fmt.Sprintf("%s-%s", nodeInfo.GetClientId(), name)
}
