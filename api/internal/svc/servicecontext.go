package svc

import (
	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/logdispatcher"
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/common/ws"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	DataKeyRedisClient datakeyredis.IF
	NodeDBClient       model.NodeModel
	RawLogDBClient     model.RawlogModel
	ListenWs           ws.IF
	LogDispatcher      logdispatcher.IF
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewSqlConn("postgres", c.PostgresConf)
	redisCli := redis.MustNewRedis(c.RedisConf)
	wsl := ws.New()
	return &ServiceContext{
		Config:             c,
		DataKeyRedisClient: datakeyredis.New(redisCli),
		NodeDBClient:       model.NewNodeModel(dbConn),
		RawLogDBClient:     model.NewRawlogModel(dbConn),
		ListenWs:           wsl,
		LogDispatcher:      logdispatcher.New(wsl),
	}
}
