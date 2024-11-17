package datakeyredis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/report"
	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	TimeLimit = time.Minute
)

type IF interface {
	SetKey(ctx context.Context, nodeInfo *report.NodeInfo, provider *report.Provider, keys []string) error
	GetKeys(ctx context.Context, nodeInfo *report.NodeInfo) ([]*DataKey, error)
}

func New(cli *redis.Redis) IF {
	return &impl{
		cli: cli,
	}
}

type impl struct {
	cli *redis.Redis
}

type DataKey struct {
	Provider *report.Provider `json:"provider"`
	Key      string           `json:"key"`
}

func (i *impl) SetKey(ctx context.Context, nodeInfo *report.NodeInfo, provider *report.Provider, keys []string) error {
	keys = lo.Map(keys, func(k string, _ int) string {
		dk := &DataKey{
			Provider: provider,
			Key:      k,
		}
		r, _ := json.Marshal(dk)
		return string(r)
	})
	nk := rediskey.LogDataKey(nodeInfo)
	now := time.Now().Format(time.DateTime)
	errs := parallel.Map(keys, func(k string, _ int) error {
		return i.cli.HsetCtx(ctx, nk, k, now)
	})
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) GetKeys(ctx context.Context, nodeInfo *report.NodeInfo) ([]*DataKey, error) {
	res, err := i.cli.HgetallCtx(ctx, rediskey.LogDataKey(nodeInfo))
	if err != nil {
		return nil, err
	}
	dataKeys := lo.MapToSlice(res, func(k string, t string) *DataKey {
		createAt := lo.Must(time.Parse(time.DateTime, t))
		if time.Since(createAt) > TimeLimit {
			go func() {
				_, _ = i.cli.HdelCtx(ctx, rediskey.LogDataKey(nodeInfo), k)
			}()
			return nil
		}
		var dk DataKey
		_ = json.Unmarshal([]byte(k), &dk)
		return &dk
	})
	dataKeys = lo.Filter(dataKeys, func(dk *DataKey, _ int) bool { return dk != nil })
	return dataKeys, nil
}
