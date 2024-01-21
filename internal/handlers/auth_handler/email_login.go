package auth_handler

import (
	msg_const "go_server/config/messages"
	"go_server/internal/db"
	"go_server/utils"

	"github.com/gofiber/fiber/v2"
)

func EmailLogin(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(c, msg_const.ErrorParsingBody, err)
	}

	user, err := db.GetUserByEmail(data["email"])
	if err != nil {
		return utils.ErrorResponse(c, msg_const.ErrorGettingUser, err)
	}

	if user == nil {
		return utils.ErrorResponse(c, msg_const.UserDoesNotExist, nil)
	}

	if !user.Verified {
		return utils.ErrorResponse(c, msg_const.UserNotVerified, nil)
	}

	if !utils.CheckPasswordHash(data["password"], user.Password) {
		return utils.ErrorResponse(c, msg_const.InvalidCredentials, nil)
	}

	token, err := utils.CreateJWT(user.ID.Hex())
	if err != nil {
		return utils.ErrorResponse(c, msg_const.ErrorGeneratingJWT, err)
	}

	return utils.SuccessResponse(c, msg_const.Success, token)
}
