package datakeyredis

import (
	"context"
	"encoding/json"

	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/report"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type IF interface {
	SetKey(ctx context.Context, nodeInfo *report.NodeInfo, provider *report.Provider, keys []string) error
}

func New(cli *redis.Redis) IF {
	return &impl{
		cli: cli,
	}
}

type impl struct {
	cli *redis.Redis
}

func (i *impl) SetKey(ctx context.Context, nodeInfo *report.NodeInfo, provider *report.Provider, keys []string) error {
	keys = lo.Map(keys, func(k string, _ int) string {
		dk := &report.DataKey{
			Provider: provider,
			Key:      k,
		}
		r, _ := json.Marshal(dk)
		return string(r)
	})
	_, err := i.cli.SaddCtx(ctx,
		rediskey.LogDataKey(nodeInfo),
		lo.ToAnySlice(keys)...)
	return err
}
