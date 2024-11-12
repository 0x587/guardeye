package svc

import (
	"github.com/0x587/guardeye/api/internal/config"
	model2 "github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/common/ws"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	NodeDBClient   model2.NodeModel
	RawLogDBClient model2.RawLogModel
	Ws             struct {
		NodeLog ws.IF
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
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
		Ws: struct{ NodeLog ws.IF }{
			NodeLog: ws.New(),
		},
	}
}
