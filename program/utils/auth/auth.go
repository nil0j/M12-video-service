package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(200 * time.Hour)
	claims["authorized"] = true
	claims["user"] = "username"

	tokenString, err := token.SignedString(config.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
