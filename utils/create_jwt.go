package utils

import (
	"errors"
	"go_server/config"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJWT(userId string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New(config.JWTNotSetMessage)
	}

	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
