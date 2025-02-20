package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/0x587/guardeye/wsapi/internal/config"
	"github.com/0x587/guardeye/wsapi/internal/ws"
)

type ServiceContext struct {
	Config config.Config
	Ws     ws.IF
	Redis  *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Ws:     ws.New(),
		Redis:  redis.MustNewRedis(c.Redis),
	}
}
