package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	CdrServiceTarget string
	PostgresConf     struct {
		Driver string
		Dsn    string
	}
}
