package main

import (
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.Init(conf)
}

func main() {
	conf := config.Get()
	logs := logger.Get("main")
	logs.Debug().Msg("Hello, World!")
	logs.Debug().Msgf("ENV: %s", conf.Env)
}
