package handlers

import (
	"go_server/config"
	"go_server/config/types"
	"go_server/internal/db"

	"go_server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// VerifyNewUser handles verification of a new user.
// It parses the request params to extract the token string, and checks if the token exists in the database.
// If the token exists, it updates the user's verified status to true.
// If successful, it returns a JSON response with the updated user's details.
//

func VerifyNewUser(c *fiber.Ctx) error {

	token := c.Params("token")
	user, err := db.GetUserByToken(token)

	if err != nil {
		return utils.ErrorResponse(c, config.UserNotFound, err)
	}

	userBSON := bson.M{
		"updatedAt": time.Now(),
		"verified":  true,
	}

	var userId = bson.M{"_id": user.ID}

	_, err = db.UpdateOne(config.Users, userId, userBSON)

	if err != nil {
		return utils.ErrorResponse(c, config.UserUpdateFailed, err)
	}

	_, err = db.DeleteOneById(config.Tokens, token)

	if err != nil {
		return utils.ErrorResponse(c, config.TokenDeleteFailed, err)
	}

	userSuccessResponse := types.UserSuccessResponse{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID.Hex(),
	}

	return utils.SuccessResponse(c, config.SuccessfulVerification, userSuccessResponse)
}
