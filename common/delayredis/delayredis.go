package delayredis

import (
	"context"
	"encoding/json"

	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/report"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type IF interface {
	SetDelay(ctx context.Context, nodeInfo *report.NodeInfo, delay *report.DelayResult) error
}

func New(cli *redis.Redis) IF {
	return &impl{
		cli: cli,
	}
}

type impl struct {
	cli *redis.Redis
}

func (i *impl) SetDelay(ctx context.Context, nodeInfo *report.NodeInfo, delay *report.DelayResult) error {
	bytes, err := json.Marshal(delay)
	if err != nil {
		return err
	}
	if err = i.cli.SetCtx(ctx, rediskey.TransDelayKey(nodeInfo), string(bytes)); err != nil {
		return err
	}
	return nil
}
