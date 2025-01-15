package model

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
)

type GoogleUserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

type AuthToken struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type AuthTokenClaims struct {
	AuthToken
	jwt.RegisteredClaims
}

type User struct {
	Id           uuid.UUID         `json:"id"`
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	Role         constant.AuthRole `json:"role"`
	IsActive     bool              `json:"is_active"`
	LastLoggedIn time.Time         `json:"last_logged_in"`
	WithTimestamp
}

type NewUser struct {
	Email string
	Name  string
}

type AuthSocialPayload struct {
	Token   *string `json:"token"`
	Message string  `json:"message"`
}

func (s *AuthSocialPayload) AddToken(token string) {
	if token == "" {
		return
	}

	s.Token = &token
}

func (s *AuthSocialPayload) RedirectUrl(conf *config.Config) string {
	res, _ := json.Marshal(s)
	token := base64.StdEncoding.EncodeToString(res)
	redirectUrl := conf.Auth.Callback
	query := redirectUrl.Query()
	query.Set("token", token)
	redirectUrl.RawQuery = query.Encode()
	return redirectUrl.String()
}
