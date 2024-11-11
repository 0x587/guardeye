package svc

import (
	"github.com/0x587/guardeye/report/internal/config"
	"github.com/0x587/guardeye/report/internal/model"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config             config.Config
	RawLogPusherClient *kq.Pusher
	RedisClient        *redis.Redis
	NodeDBClient       model.NodeModel
	RawLogDBClient     model.RawLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RawLogPusherClient: kq.NewPusher(
			c.RawLogPusherConf.Brokers,
			c.RawLogPusherConf.Topic,
		),
		RedisClient: redis.MustNewRedis(c.ReportRedis),
		NodeDBClient: model.NewNodeModel(
			c.MongoConf.Uri,
			c.MongoConf.Database,
			"node",
		),
		RawLogDBClient: model.NewRawLogModel(
			c.MongoConf.Uri,
			c.MongoConf.Database,
			"rawlog",
		),
	}
}
