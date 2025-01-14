package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/handler"
)

type Router interface {
	HomeRoute(handler handler.HomeHandler)
}

type router struct {
	app *fiber.App
}

func NewRouter(app *fiber.App) Router {
	return &router{
		app: app,
	}
}

func (r *router) HomeRoute(handler handler.HomeHandler) {
	r.app.Get("", handler.Home)
}
