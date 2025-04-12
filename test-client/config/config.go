package config

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server       string
	LinkEndpoint string
	Provider     []ProviderConfig
	Bridge       BridgeConfig
}
type ProviderConfig struct {
	Type string
	Args map[string]interface{}
}

type BridgeConfig struct {
	Enable bool
	Ip     string
	Port   int
}

var configFile = flag.String("f", "/etc/guardeye-agent/config.yaml", "the config file")

func LoadConfig() Config {
	var c Config
	conf.MustLoad(*configFile, &c)
	return c
}
