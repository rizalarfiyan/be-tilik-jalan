package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
	"github.com/rizalarfiyan/be-tilik-jalan/constant"
	"github.com/rizalarfiyan/be-tilik-jalan/exception"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/model"
	"github.com/rizalarfiyan/be-tilik-jalan/internal/repository"
	"github.com/rizalarfiyan/be-tilik-jalan/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthService interface {
	Google(state ...oauth2.AuthCodeOption) string
	GoogleCallback(ctx context.Context, code string) (redirectUrl string)
}

type authService struct {
	conf      *config.Config
	exception exception.Exception
	google    oauth2.Config
	repo      repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	conf := config.Get()
	return &authService{
		conf:      conf,
		exception: exception.NewException(),
		google: oauth2.Config{
			ClientID:     conf.Auth.Google.ClientID,
			ClientSecret: conf.Auth.Google.ClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  conf.Auth.Google.RedirectURL,
			Scopes:       conf.Auth.Google.Scopes,
		},
		repo: repo,
	}
}

func (s *authService) Google(state ...oauth2.AuthCodeOption) string {
	return s.google.AuthCodeURL(constant.AUTH_STATE, state...)
}

func (s *authService) GoogleCallback(ctx context.Context, code string) (redirectUrl string) {
	var token string
	defer func() {
		var message string
		if r := recover(); r != nil {
			message = fmt.Sprint(r)
		} else {
			message = "Successfully login with google"
		}

		data := model.AuthSocialPayload{
			Message: message,
		}
		data.AddToken(token)
		redirectUrl = data.RedirectUrl(s.conf)
	}()

	code = strings.TrimSpace(code)
	s.exception.UnauthorizedBool(code == "", "Invalid oauth2 google code")

	googleToken, err := s.google.Exchange(ctx, code)
	s.exception.UnauthorizedErr(err, "Invalid oauth2 google code")

	client := &http.Client{
		Timeout: constant.DEFAULT_API_TIMEOUT,
	}
	endpoint := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?access_token=%s", url.QueryEscape(googleToken.AccessToken))
	ctx, cancel := context.WithTimeout(ctx, constant.DEFAULT_API_TIMEOUT)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	s.exception.UnauthorizedErr(err, "Invalid get user info oauth2 google request")

	resp, err := client.Do(req)
	s.exception.UnauthorizedErr(err, "Invalid get user info oauth2 google response")
	defer resp.Body.Close()

	var userInfo model.GoogleUserInfo
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	s.exception.UnauthorizedErr(err, "Invalid decode user info oauth2 google response")

	authToken := s.getOrRegisterUser(ctx, userInfo)
	token, err = s.createToken(authToken)
	s.exception.Error(err)
	return
}

func (s *authService) getOrRegisterUser(ctx context.Context, userInfo model.GoogleUserInfo) model.AuthToken {
	info := model.AuthToken{
		Email: userInfo.Email,
	}

	user, err := s.repo.GetByEmail(ctx, userInfo.Email)
	s.exception.ErrorSkipNotFound(err)
	if user != nil {
		info.Id = user.Id
		return info
	}

	id, err := s.repo.Insert(ctx, model.NewUser{
		Email: userInfo.Email,
		Name:  userInfo.Name,
	})
	s.exception.Error(err)

	info.Id = *id
	return info
}

func (s *authService) createToken(user model.AuthToken) (string, error) {
	now := time.Now()
	claims := model.AuthTokenClaims{
		AuthToken: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprint(user.Id),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.conf.Auth.JWT.Expired)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	return utils.GenerateJwtToken(claims, s.conf.Auth.JWT.Secret)
}
