package handlers

import (
	"go_server/config"
	"go_server/utils"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(c *fiber.Ctx) error {
	secret := os.Getenv("JWT_SECRET")
	jwtToken := c.Params("jwt_token")

	if secret == "" {
		return utils.ErrorResponse(c, config.JWTNotSetMessage, nil)
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return utils.ErrorResponse(c, config.ErrorParsingJWT, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return utils.ErrorResponse(c, config.InvalidJWT, nil)
	}

	userId := claims["userId"].(string)

	return utils.SuccessResponse(c, config.Success, userId)

}
