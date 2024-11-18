package mqs

import (
	"context"

	"github.com/0x587/guardeye/report/internal/config"
	"github.com/0x587/guardeye/report/internal/mqs/logtodb"
	"github.com/0x587/guardeye/report/internal/mqs/logtometric"
	"github.com/0x587/guardeye/report/internal/mqs/parselogkey"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.LogtodbConsumerConf, logtodb.New(ctx, svcContext)),
		kq.MustNewQueue(c.LogDataKeyConsumerConf, parselogkey.New(ctx, svcContext)),
		kq.MustNewQueue(c.LogToMetricConsumerConf, logtometric.New(ctx, svcContext)),
	}
}
