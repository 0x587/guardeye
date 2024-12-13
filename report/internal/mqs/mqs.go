package mqs

import (
	"context"

	"github.com/zeromicro/go-zero/core/service"

	"github.com/0x587/guardeye/report/internal/config"
	"github.com/0x587/guardeye/report/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		//kq.MustNewQueue(c.LogParseConsumerConf, logparse.New(ctx, svcContext)),
	}
}
