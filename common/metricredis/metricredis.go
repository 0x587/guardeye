package metricredis

import (
	"context"
	"fmt"
	"time"

	"github.com/0x587/guardeye/report/report"
	"github.com/redis/go-redis/v9"
)

type IF interface {
	AddMetric(ctx context.Context, nodeInfo *report.NodeInfo, name string, value float64) error
}

func New(cli *redis.Client) IF {
	return &impl{
		cli: cli,
	}
}

type impl struct {
	cli *redis.Client
}

func (i *impl) AddMetric(ctx context.Context, nodeInfo *report.NodeInfo, name string, value float64) error {
	cmd := i.cli.TSAdd(ctx, fmt.Sprintf("%s-%s", nodeInfo.GetClientId(), name), time.Now().UnixMilli(), value)
	return cmd.Err()
}
