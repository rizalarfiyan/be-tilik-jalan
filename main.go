package main

import (
	"context"

	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/database"
	"github.com/rizalarfiyan/be-tilik-jalan/internal"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rizalarfiyan/be-tilik-jalan/validation"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.Init(conf)

	ctx := context.Background()
	database.InitPostgresql(ctx, conf)
	validation.Init()
}

// @title						BE Tilik Jalan
// @version					1.0
// @description				This is an API documentation of BE Tilik Jalan
// @BasePath					/
// @securityDefinitions.apikey	AccessToken
// @in							header
// @name						Authorization
func main() {
	internal.Run()
}
