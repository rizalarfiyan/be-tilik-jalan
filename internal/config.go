package internal

import (
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/exception"
	"github.com/rs/zerolog"
)

func fiberConfig(conf *config.Config) fiber.Config {
	return fiber.Config{
		ErrorHandler:                 exception.ErrorHandler,
		BodyLimit:                    conf.Http.MaxUploadSize * 1024 * 1024,
		StreamRequestBody:            true,
		DisablePreParseMultipartForm: true,
	}
}

func corsConfig(conf *config.Config) cors.Config {
	return cors.Config{
		AllowOrigins:     conf.Cors.AllowOrigins,
		AllowMethods:     conf.Cors.AllowMethods,
		AllowHeaders:     conf.Cors.AllowHeaders,
		AllowCredentials: conf.Cors.AllowCredentials,
		ExposeHeaders:    conf.Cors.ExposeHeaders,
	}
}

func zerologConfig(logApi *zerolog.Logger) fiberzerolog.Config {
	fields := []string{
		fiberzerolog.FieldIP,
		fiberzerolog.FieldMethod,
		fiberzerolog.FieldPath,
		fiberzerolog.FieldURL,
		fiberzerolog.FieldMethod,
		fiberzerolog.FieldPath,
		fiberzerolog.FieldLatency,
		fiberzerolog.FieldStatus,
		fiberzerolog.FieldBody,
		fiberzerolog.FieldError,
		fiberzerolog.FieldRequestID,
	}

	return fiberzerolog.Config{
		Logger: logApi,
		Fields: fields,
		SkipBody: func(ctx *fiber.Ctx) bool {
			return strings.Contains(string(ctx.Request().Header.ContentType()), fiber.MIMEMultipartForm)
		},
	}
}
