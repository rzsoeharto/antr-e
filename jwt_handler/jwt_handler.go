package jwthandler

import (
	logger "inline/Logger"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken() string {
	key := []byte(os.Getenv("SECRET_KEY"))
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7)),
		Issuer:    "Antr-e",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := token.SignedString(key)

	if err != nil {
		logger.Logger.Panicln(err)
		return ""
	}

	return SignedToken
}
