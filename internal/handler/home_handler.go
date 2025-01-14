package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/response"
)

type HomeHandler interface {
	Home(c *fiber.Ctx) error
}

type homeHandler struct {
}

func NewHomeHandler() HomeHandler {
	return &homeHandler{}
}

// Home godoc
//
//	@Summary		Home based on parameter
//	@Description	Home
//	@ID				home
//	@Tags			Home
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Base{data=response.Home}
//	@Failure		500	{object}	response.Base
//	@Router			/ [get]
func (h *homeHandler) Home(c *fiber.Ctx) error {
	return c.JSON(response.Base{
		Code:    fiber.StatusOK,
		Message: "Success!",
		Data: response.Home{
			Title:    "API Tilik Jalan",
			Author:   "Muhamad Rizal Arfiyan",
			Github:   "https://github.com/rizalarfiyan/",
			Linkedin: "https://www.linkedin.com/in/rizalarfiyan/",
			Website:  "https://rizalarfiyan.com/",
		},
	})
}
