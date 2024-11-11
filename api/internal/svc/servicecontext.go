package svc

import (
	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/report/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	NodeDBClient   model.NodeModel
	RawLogDBClient model.RawLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
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
