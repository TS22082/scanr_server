package handlers

import (
	"bytes"
	db_const "go_server/config/db"
	msg_const "go_server/config/messages"
	"go_server/config/types"
	"go_server/internal/db"
	"go_server/internal/models"
	"go_server/utils"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateNewUser handles the creation of a new user.
// It parses the request body to extract user details, hashes the user's password,
// and inserts the new user record into the database.
// If successful, it returns a JSON response with the created user's details.
//
// The function responds with an error message and appropriate status code in the following cases:
// - Failure to parse the request body.
// - Failure to hash the password.
// - Failure to insert the new user into the database.
// - Failure to assert the type of the inserted ObjectID.
// - Failure to insert the new register token into the database.
// - Failure to assert the type of the inserted ObjectID.
// - Failure to send the verification email.
//
// Parameters:
// - c: Fiber context which includes the HTTP request data.
//
// Returns:
// - Returns a Fiber error, if any, during the processing of the request.

func CreateNewUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(c, msg_const.ErrorParsingBody, err)
	}

	password, err := utils.HashPassword(data["password"])
	if err != nil {
		return utils.ErrorResponse(c, msg_const.PasswordHashingError, err)
	}

	newUser := &models.User{
		Username:  data["username"],
		Email:     data["email"],
		Verified:  false,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userAlreadyExists, err := db.UserExists(newUser.Email)
	if err != nil {
		// Handle other errors (e.g., database connection errors)
		return utils.ErrorResponse(c, msg_const.ErrorCheckingUserExists, err)
	}

	if userAlreadyExists {
		// Handle the case where the user already exists
		return utils.ErrorResponse(c, msg_const.UserExistsError, nil)
	}

	result, err := db.Create(db_const.Users, newUser)

	if err != nil {
		return utils.ErrorResponse(c, msg_const.UserCreateFailed, err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return utils.ErrorResponse(c, msg_const.ObjectAssertionError, nil)
	}

	newUser.ID = oid

	newRegisterToken := &models.Register_Token{
		UserID:    newUser.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err = db.Create(db_const.Tokens, newRegisterToken)
	if err != nil {
		return utils.ErrorResponse(c, msg_const.TokenCreatedFailed, err)
	}

	oid, ok = result.InsertedID.(primitive.ObjectID)
	if !ok {
		return utils.ErrorResponse(c, msg_const.ObjectAssertionError, nil)
	}

	newRegisterToken.ID = oid

	domain := os.Getenv("CLIENT_DOMAIN")

	templateData := types.TemplateData{
		Domain: domain,
		Token:  newRegisterToken.ID.Hex(),
	}

	templatePath := filepath.Join("templates", "verification_template.html")

	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return utils.ErrorResponse(c, msg_const.TemplateParseError, err)
	}

	var emailBody bytes.Buffer

	if err := template.Execute(&emailBody, templateData); err != nil {
		return utils.ErrorResponse(c, msg_const.TemplateExecError, err)
	}

	if err := utils.SendEmail(newUser.Email, msg_const.VerifyEmailSubject, emailBody.String()); err != nil {
		return utils.ErrorResponse(c, msg_const.EmailSendError, err)
	}

	userSuccessResponse := types.UserSuccessResponse{
		Username: newUser.Username,
		Email:    newUser.Email,
		ID:       newUser.ID.Hex(),
	}

	return utils.SuccessResponse(c, msg_const.UserCreateSuccess, userSuccessResponse)
}
