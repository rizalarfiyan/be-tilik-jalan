package internal

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/database"
	_ "github.com/rizalarfiyan/be-tilik-jalan/docs"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/handler"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/repository"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/service"
	"github.com/rizalarfiyan/be-tilik-jalan/logger"
	"github.com/rizalarfiyan/be-tilik-jalan/middleware"
	"github.com/rs/zerolog"
)

func Run() {
	logs := logger.Get("app")
	logs.Info().Msg("Application is running!")

	pgSql := database.GetPostgresql()
	defer func(pgSql *sql.DB) {
		err := pgSql.Close()
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to close database connection")
		}
	}(pgSql)

	conf := config.Get()
	logs.Info().Msg("Configuring server...")
	logApi := logger.Get("api")
	app := fiber.New(fiberConfig(conf))
	app.Use(fiberzerolog.New(zerologConfig(logApi)))
	app.Use(requestid.New())
	app.Use(cors.New(corsConfig(conf)))
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(recover.New())

	logs.Info().Msg("Server is running!")
	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	server := &http.Server{
		Addr: baseUrl,
	}

	app.Get("/swagger/*", basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.Swagger.Username: conf.Swagger.Password,
		},
	}), swagger.New(swagger.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  true,
		DocExpansion: "list",
	}))

	go func() {
		err := app.Listen(baseUrl)
		if err != nil {
			logs.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Repository
	authRepository := repository.NewAuthRepository(pgSql)

	// Service
	authService := service.NewAuthService(authRepository)

	// Handler
	homeHandler := handler.NewHomeHandler()
	authHandler := handler.NewAuthHandler(authService)

	// Router
	mid := middleware.NewMiddleware(authRepository)
	router := NewRouter(app, mid)
	router.HomeRoute(homeHandler)
	router.AuthRoute(authHandler)

	handleShutdown(server, logs)
}

func handleShutdown(server *http.Server, logs *zerolog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Warn().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = server.Shutdown(ctx); err != nil {
		logs.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logs.Info().Msg("Server exiting")
}
