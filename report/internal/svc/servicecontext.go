package svc

import (
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/delayredis"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/internal/config"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	RawLogPusherClient  *kq.Pusher
	DataKeyRedisClient  datakeyredis.IF
	DelayRedisClient    delayredis.IF
	MetricRedisClient   metricredis.IF
	NodeDBClient        model.NodeModel
	RawLogDBClient      model.RawlogModel
	LogToMetricDBClient model.LogToMetricQueryModel
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
		DataKeyRedisClient:  datakeyredis.New(redisCli),
		DelayRedisClient:    delayredis.New(redisCli),
		MetricRedisClient:   metricredis.New(goredis.NewClient(&goredis.Options{Addr: c.ReportRedis.Host})),
		NodeDBClient:        model.NewNodeModel(dbConn),
		RawLogDBClient:      model.NewRawlogModel(dbConn),
		LogToMetricDBClient: model.NewLogToMetricQueryModel(dbConn),
	}
}
