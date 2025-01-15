package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/response"
)

func (m *middleware) Auth(permissions ...constant.AuthRole) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authTokens, isFound := ctx.GetReqHeaders()[fiber.HeaderAuthorization]
		if !isFound || len(authTokens) == 0 {
			return m.errUnauthorized(ctx)
		}

		authToken := authTokens[0]
		if len(authToken) < 7 || authToken[:7] != "Bearer " {
			return m.errUnauthorized(ctx)
		}

		authToken = authToken[7:]
		payload, isValid := m.validateJwtToken(authToken)
		if !isValid {
			return m.errUnauthorized(ctx)
		}

		user, err := m.authRepo.GetById(ctx.Context(), payload.Id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return m.errForbidden(ctx)
			}
			return m.errUnauthorized(ctx)
		}

		if !user.IsActive {
			return m.errForbidden(ctx)
		}

		if isValid := user.Role.Have(permissions...); !isValid {
			return m.errForbidden(ctx)
		}

		ctx.Locals(constant.AUTH_KEY_LOCALS, *user)
		return ctx.Next()
	}
}

func (m *middleware) errUnauthorized(ctx *fiber.Ctx) error {
	code := fiber.StatusUnauthorized
	return ctx.Status(code).JSON(response.Base{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	})
}

func (m *middleware) errForbidden(ctx *fiber.Ctx) error {
	code := fiber.StatusForbidden
	return ctx.Status(code).JSON(response.Base{
		Code:    code,
		Message: http.StatusText(code),
		Data:    nil,
	})
}

func (m *middleware) validateJwtToken(accessToken string) (*model.AuthTokenClaims, bool) {
	if len(accessToken) == 0 {
		return nil, false
	}

	var tokenClaims model.AuthTokenClaims
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.conf.Auth.JWT.Secret), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil || !token.Valid {
		return nil, false
	}

	if _, ok := token.Claims.(*model.AuthTokenClaims); !ok {
		return nil, false
	}

	return &tokenClaims, true
}
