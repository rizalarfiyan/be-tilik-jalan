package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/handler"
	"github.com/rizalarfiyan/be-tilik-jalan/middleware"
)

type Router interface {
	HomeRoute(handler handler.HomeHandler)
	AuthRoute(handler handler.AuthHandler)
}

type router struct {
	app *fiber.App
	mid middleware.Middleware
}

func NewRouter(app *fiber.App, mid middleware.Middleware) Router {
	return &router{
		app: app,
		mid: mid,
	}
}

func (r *router) HomeRoute(handler handler.HomeHandler) {
	r.app.Get("", handler.Home)
}

func (r *router) AuthRoute(handler handler.AuthHandler) {
	r.app.Get("/auth/me", r.mid.Auth(), handler.Me)
	r.app.Get("/auth/google", handler.Google)
	r.app.Get("/auth/google/callback", handler.GoogleCallback)
}
