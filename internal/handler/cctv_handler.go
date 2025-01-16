package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/response"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/service"
)

type CCTVHandler interface {
	GetAll(ctx *fiber.Ctx) error
}

type cctvHandler struct {
	service service.CCTVService
}

func NewCCTVHandler(service service.CCTVService) CCTVHandler {
	return &cctvHandler{
		service: service,
	}
}

// GetAll godoc.
//
//	@Summary	Get all CCTV records
//	@ID			get-all-cctv
//	@Tags		cctv
//	@Accept		json
//	@Produce	json
//	@Security	AccessToken
//	@Success	200	{object}	response.Base{data=model.CCTVs}
//	@Failure	500	{object}	response.Base{message=string}
//	@Router		/cctv [get]
func (h *cctvHandler) GetAll(ctx *fiber.Ctx) error {
	list := h.service.GetAll(ctx.Context())
	return ctx.JSON(response.Base{
		Code:    fiber.StatusOK,
		Message: "Get all cctv successfully",
		Data:    list,
	})
}
