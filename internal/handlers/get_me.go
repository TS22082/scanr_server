package handlers

import (
	"go_server/config"
	"go_server/config/types"
	"go_server/internal/db"
	"go_server/utils"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {

	userId := c.Locals("userId").(string)
	user, err := db.GetUserById(userId)

	if err != nil {
		return utils.ErrorResponse(c, config.ErrorGettingUser, err)
	}

	responseWithUser := types.UserSuccessResponse{
		Username: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		ID:       user.ID.Hex(),
	}

	return utils.SuccessResponse(c, config.Success, responseWithUser)
}
