package handlers

import (
	"go_server/config"
	"go_server/internal/db"
	"go_server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ResetPassword(c *fiber.Ctx) error {
	var data = map[string]string{}

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(c, config.ErrorParsingBody, err)
	}

	token := c.Params("token")
	user, err := db.GetUserByToken(token)

	if err != nil {
		return utils.ErrorResponse(c, config.UserNotFound, err)
	}

	newHashedPassword, err := utils.HashPassword(data["password"])

	if err != nil {
		return utils.ErrorResponse(c, config.PasswordHashingError, err)
	}

	userBSON := bson.M{
		"updatedAt": time.Now(),
		"password":  newHashedPassword,
	}

	var userId = bson.M{"_id": user.ID}

	_, err = db.UpdateOne(config.Users, userId, userBSON)

	if err != nil {
		return utils.ErrorResponse(c, config.UserNotFound, err)
	}

	_, err = db.DeleteOneById(config.Tokens, token)

	if err != nil {
		return utils.ErrorResponse(c, config.TokenDeleteFailed, err)
	}

	return utils.SuccessResponse(c, config.Success, nil)

}
