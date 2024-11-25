package svc

import (
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/delayredis"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/common/svcctx"
	"github.com/0x587/guardeye/report/internal/config"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	svcctx.ServiceContext
	Config             config.Config
	RawLogPusherClient *kq.Pusher
	DataKeyRedisClient datakeyredis.IF
	DelayRedisClient   delayredis.IF
	MetricRedisClient  metricredis.IF
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
		MetricRedisClient:  metricredis.New(goredis.NewClient(&goredis.Options{Addr: c.ReportRedis.Host})),
		ServiceContext:     *svcctx.NewServiceContext(dbConn),
	}
}
