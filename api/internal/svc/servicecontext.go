package svc

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	goredis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/logdispatcher"
	"github.com/0x587/guardeye/api/internal/ws"
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/common/svcctx"
)

type ServiceContext struct {
	svcctx.ServiceContext
	Config             config.Config
	DataKeyRedisClient datakeyredis.IF
	ListenWs           ws.IF
	MetricRedisClient  metricredis.IF
	LogDispatcher      logdispatcher.IF
	Es                 *elasticsearch.Client
	Minio              *minio.Client
	Redis              *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewSqlConn("postgres", c.PostgresConf)
	redisCli := redis.MustNewRedis(c.RedisConf)
	wsl := ws.New()
	es := lo.Must(elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://ws.scut.mcurobot.com:59200"},
	}))
	minioCli := lo.Must(minio.New("minioapi.guardeye.shawnsiu.site:5080", &minio.Options{
		Creds: credentials.NewStaticV4("fJGShizzQOmRu44EVfgv", "47kDXUYY8dusuqMZCCsHCQQANB3l5WpXTMczh580", ""),
	}))
	return &ServiceContext{
		Config:             c,
		DataKeyRedisClient: datakeyredis.New(redisCli),
		ListenWs:           wsl,
		LogDispatcher:      logdispatcher.New(wsl),
		MetricRedisClient:  metricredis.New(goredis.NewClient(&goredis.Options{Addr: c.RedisConf.Host})),
		ServiceContext:     *svcctx.NewServiceContext(dbConn),
		Es:                 es,
		Minio:              minioCli,
		Redis:              redisCli,
	}
}
