package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(claims jwt.Claims, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetGravatar(email string) string {
	hasher := sha256.Sum256([]byte(strings.TrimSpace(email)))
	hash := hex.EncodeToString(hasher[:])
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=50&r=pg&d=404", hash)
}
