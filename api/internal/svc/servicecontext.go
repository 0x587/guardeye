package svc

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	goredis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/logdispatcher"
	"github.com/0x587/guardeye/api/internal/ws"
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/common/svcctx"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
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
	//Mqtt               mqtt.Client
	CdrCli foxgloveService.FoxgloveServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbConn := sqlx.NewSqlConn("postgres", c.PostgresConf)
	redisCli := redis.MustNewRedis(c.RedisConf)
	wsl := ws.New()
	es := lo.Must(elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://ws.scut.mcurobot.com:59200"},
	}))
	minioCli := lo.Must(minio.New(c.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(c.Minio.AccessKey, c.Minio.AccessSecret, ""),
	}))

	cdrClient := lo.Must(grpc.NewClient("127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())))
	cdrCli := foxgloveService.NewFoxgloveServiceClient(cdrClient)

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
		//Mqtt:               mqttCli,
		CdrCli: cdrCli,
	}
}
