package handlers

import (
	"go_server/config"
	"go_server/config/types"
	"go_server/internal/db"
	"go_server/utils"

	"github.com/gofiber/fiber/v2"
)

func EmailLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(c, config.ErrorParsingBody, err)
	}

	user, err := db.GetUserByEmail(data["email"])

	if err != nil {
		return utils.ErrorResponse(c, config.ErrorGettingUser, err)
	}

	if user == nil {
		return utils.ErrorResponse(c, config.UserDoesNotExist, nil)
	}

	if !user.Verified {
		return utils.ErrorResponse(c, config.UserNotVerified, nil)
	}

	if !utils.CheckPasswordHash(data["password"], user.Password) {
		return utils.ErrorResponse(c, config.InvalidCredentials, nil)
	}

	token, err := utils.CreateJWT(user.ID.Hex())
	if err != nil {
		return utils.ErrorResponse(c, config.ErrorGeneratingJWT, err)
	}

	LoginUserSuccessResponse := &types.LoginUserSuccessResponse{
		Email:    user.Email,
		Username: user.Username,
		ID:       user.ID.Hex(),
		Token:    token,
	}

	return utils.SuccessResponse(c, config.Success, LoginUserSuccessResponse)
}
