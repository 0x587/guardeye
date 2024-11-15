package mqs

import (
	"context"

	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/mqs/hotlog"
	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.HotLogConsumerConf, hotlog.New(ctx, svcContext)),
	}
}
