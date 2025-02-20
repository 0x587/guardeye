package mqs

import (
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/0x587/guardeye/wsapi/internal/config"
	"github.com/0x587/guardeye/wsapi/internal/mqs/sendtows"
	"github.com/0x587/guardeye/wsapi/internal/svc"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.WsSubscribeConsumerConf, sendtows.New(ctx, svcContext)),
	}
}
