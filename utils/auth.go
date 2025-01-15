package utils

import "github.com/golang-jwt/jwt/v5"

func GenerateJwtToken(claims jwt.Claims, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
