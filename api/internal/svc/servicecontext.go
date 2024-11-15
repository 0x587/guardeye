package svc

import (
	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/common/ws"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Redis
	NodeDBClient   model.NodeModel
	RawLogDBClient model.RawlogModel
	BoardCaseWs    map[string]ws.IF[string]
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewSqlConn("postgres", c.PostgresConf)
	return &ServiceContext{
		Config:         c,
		RedisClient:    redis.MustNewRedis(c.RedisConf),
		NodeDBClient:   model.NewNodeModel(dbConn),
		RawLogDBClient: model.NewRawlogModel(dbConn),
		BoardCaseWs:    make(map[string]ws.IF[string]),
	}
}
