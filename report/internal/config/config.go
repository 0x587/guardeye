package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RawLogPusherConf struct {
		Brokers []string
		Topic   string
	}
	ReportRedis redis.RedisConf
	MongoConf   struct {
		Uri      string
		Database string
	}
}
