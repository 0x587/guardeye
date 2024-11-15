package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	LogtodbConsumerConf    kq.KqConf
	LogDataKeyConsumerConf kq.KqConf
	ReportRedis            redis.RedisConf
	RawLogPusherConf       struct {
		Brokers []string
		Topic   string
	}
	PostgresConf string
}
