package middleware

import (
	"go_server/config"
	"go_server/utils"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CheckJWT(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	tokenFromHeader := strings.TrimPrefix(authHeader, "Bearer ")

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return utils.ErrorResponse(c, config.JWTNotSetMessage, nil)
	}

	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return utils.ErrorResponse(c, config.ErrorParsingJWT, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return utils.ErrorResponse(c, config.InvalidJWT, nil)
	}

	userId, ok := claims["userId"].(string)
	if !ok {
		return utils.ErrorResponse(c, "User ID not found in token", nil)
	}

	c.Locals("userId", userId)
	return c.Next()
}
