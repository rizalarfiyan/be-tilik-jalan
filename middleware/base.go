package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/repository"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rs/zerolog"
)

type Middleware interface {
	Auth(roles ...constant.AuthRole) fiber.Handler
}

type middleware struct {
	authRepo repository.AuthRepository
	conf     *config.Config
	log      *zerolog.Logger
}

func NewMiddleware(authRepo repository.AuthRepository) Middleware {
	return &middleware{
		authRepo: authRepo,
		conf:     config.Get(),
		log:      logger.Get("middleware"),
	}
}
