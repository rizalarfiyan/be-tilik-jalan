package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/exception"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/response"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/service"
)

type AuthHandler interface {
	Google(ctx *fiber.Ctx) error
	GoogleCallback(ctx *fiber.Ctx) error
	Me(ctx *fiber.Ctx) error
}

type authHandler struct {
	service   service.AuthService
	conf      *config.Config
	exception exception.Exception
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		service:   service,
		conf:      config.Get(),
		exception: exception.NewException(),
	}
}

// Auth Google Redirection godoc
//
//	@Summary		Get Auth Google Redirection based on parameter
//	@Description	Auth Google Redirection
//	@ID				get-auth-google
//	@Tags			auth
//	@Success		307
//	@Failure		500
//	@Router			/auth/google [get]
func (h *authHandler) Google(ctx *fiber.Ctx) error {
	url := h.service.Google()
	return ctx.Redirect(url, http.StatusTemporaryRedirect)
}

// Auth Google Callback godoc
//
//	@Summary		Get Auth Google Callback based on parameter
//	@Description	Auth Google Callback
//	@ID				get-auth-google-callback
//	@Tags			auth
//	@Success		307
//	@Failure		500
//	@Router			/auth/google/callback [get]
func (h *authHandler) GoogleCallback(ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	redirectUrl := h.service.GoogleCallback(ctx.Context(), code)
	return ctx.Redirect(redirectUrl, http.StatusTemporaryRedirect)
}

// Auth Me godoc
//
//	@Summary		Get Auth Me based on parameter
//	@Description	Auth Me
//	@ID				get-auth-me
//	@Security		AccessToken
//	@Tags			auth
//	@Success		200	{object}	response.Base{data=response.AuthMe}
//	@Failure		500	{object}	response.Base{message=string}
//	@Router			/auth/me [get]
func (h *authHandler) Me(ctx *fiber.Ctx) error {
	user := new(model.User)
	isFound := user.FromReq(ctx)
	h.exception.UnauthorizedBool(!isFound)

	data := new(response.AuthMe)
	data.FromUser(user)
	return ctx.JSON(response.Base{
		Code:    fiber.StatusOK,
		Message: "Get auth me successfully",
		Data:    data,
	})
}
