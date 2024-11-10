package svc

import (
	"github.com/0x587/guardeye/report/internal/config"
	model "github.com/0x587/guardeye/report/internal/models"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config             config.Config
	RawLogPusherClient *kq.Pusher
	RedisClient        *redis.Redis
	NodeClient         model.NodeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RawLogPusherClient: kq.NewPusher(
			c.RawLogPusherConf.Brokers,
			c.RawLogPusherConf.Topic,
		),
		RedisClient: redis.MustNewRedis(c.ReportRedis),
		NodeClient: model.NewNodeModel(
			c.MongoConf.Uri,
			c.MongoConf.Database,
			"node",
		),
	}
}
