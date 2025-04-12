package svc

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	goredis "github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/lib/pq"

	"github.com/0x587/guardeye/api/internal/config"
	"github.com/0x587/guardeye/api/internal/logdispatcher"
	"github.com/0x587/guardeye/api/internal/ws"
	"github.com/0x587/guardeye/common/datakeyredis"
	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/metricredis"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
	"github.com/0x587/guardeye/link/linkclient"
)

type ServiceContext struct {
	Config             config.Config
	Db                 *ent.Client
	DataKeyRedisClient datakeyredis.IF
	ListenWs           ws.IF
	MetricRedisClient  metricredis.IF
	LogDispatcher      logdispatcher.IF
	Es                 *elasticsearch.Client
	Minio              *minio.Client
	Redis              *redis.Redis
	CdrCli             foxgloveService.FoxgloveServiceClient
	LinkCli            linkclient.Link
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := lo.Must(ent.Open("postgres", c.PostgresConf))
	logx.Must(db.Schema.Create(context.Background()))
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
	cli := lo.Must(zrpc.NewClient(c.Link))
	return &ServiceContext{
		Db:                 db,
		Config:             c,
		DataKeyRedisClient: datakeyredis.New(redisCli),
		ListenWs:           wsl,
		LogDispatcher:      logdispatcher.New(wsl),
		MetricRedisClient:  metricredis.New(goredis.NewClient(&goredis.Options{Addr: c.RedisConf.Host})),
		Es:                 es,
		Minio:              minioCli,
		Redis:              redisCli,
		LinkCli:            linkclient.NewLink(cli),
		CdrCli:             cdrCli,
	}
}
