package main

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Server struct {
		Ip   string
		Port int
	}
	Provider []struct {
		Type string
		Args map[string]interface{}
	}
}

var configFile = flag.String("f", "config.yaml", "the config file")

func LoadConfig() Config {
	var c Config
	conf.MustLoad(*configFile, &c)
	return c
}
