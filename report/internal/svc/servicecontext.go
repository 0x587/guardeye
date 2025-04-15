package svc

import (
	"context"

	_ "github.com/lib/pq"
	goredis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/delayredis"
	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/httpcb"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/report/internal/config"
)

type ServiceContext struct {
	Config                  config.Config
	RawLogPusherClient      *kq.Pusher
	WsSubscribePusherClient *kq.Pusher
	RedisClient             *redis.Redis
	DataKeyRedisClient      datakeyredis.IF
	DelayRedisClient        delayredis.IF
	MetricRedisClient       metricredis.IF
	Db                      *ent.Client
	HttpCallback            httpcb.IF
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCli := redis.MustNewRedis(c.ReportRedis)
	db := lo.Must(ent.Open(c.PostgresConf.Driver, c.PostgresConf.Dsn))
	logx.Must(db.Schema.Create(context.Background()))
	httpCallback := httpcb.New(db)
	return &ServiceContext{
		Config: c,
		RawLogPusherClient: kq.NewPusher(
			c.RawLogPusherConf.Brokers,
			c.RawLogPusherConf.Topic,
		),
		WsSubscribePusherClient: kq.NewPusher(
			c.WsSubscribePusherConf.Brokers,
			c.WsSubscribePusherConf.Topic,
		),
		RedisClient:        redisCli,
		DataKeyRedisClient: datakeyredis.New(redisCli),
		DelayRedisClient:   delayredis.New(redisCli),
		MetricRedisClient:  metricredis.New(goredis.NewClient(&goredis.Options{Addr: c.ReportRedis.Host})),
		Db:                 db,
		HttpCallback:       httpCallback,
	}
}
