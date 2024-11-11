package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisConf redis.RedisConf
	MongoConf struct {
		Uri      string
		Database string
	}
}
