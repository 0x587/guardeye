package svc

import (
	"github.com/0x587/guardeye/report/internal/config"
	model2 "github.com/0x587/guardeye/report/model"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config             config.Config
	RawLogPusherClient *kq.Pusher
	RedisClient        *redis.Redis
	NodeDBClient       model2.NodeModel
	RawLogDBClient     model2.RawLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RawLogPusherClient: kq.NewPusher(
			c.RawLogPusherConf.Brokers,
			c.RawLogPusherConf.Topic,
		),
		RedisClient: redis.MustNewRedis(c.ReportRedis),
		NodeDBClient: model2.NewNodeModel(
			c.MongoConf.Uri,
			c.MongoConf.Database,
			"node",
		),
		RawLogDBClient: model2.NewRawLogModel(
			c.MongoConf.Uri,
			c.MongoConf.Database,
			"rawlog",
		),
	}
}
