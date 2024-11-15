package svc

import (
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/delayredis"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/internal/config"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	RawLogPusherClient *kq.Pusher
	DataKeyRedisClient datakeyredis.IF
	DelayRedisClient   delayredis.IF
	NodeDBClient       model.NodeModel
	RawLogDBClient     model.RawlogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewSqlConn("postgres", c.PostgresConf)
	redisCli := redis.MustNewRedis(c.ReportRedis)
	return &ServiceContext{
		Config: c,
		RawLogPusherClient: kq.NewPusher(
			c.RawLogPusherConf.Brokers,
			c.RawLogPusherConf.Topic,
		),
		DataKeyRedisClient: datakeyredis.New(redisCli),
		DelayRedisClient:   delayredis.New(redisCli),
		NodeDBClient:       model.NewNodeModel(dbConn),
		RawLogDBClient:     model.NewRawlogModel(dbConn),
	}
}
