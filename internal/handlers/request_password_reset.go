package handlers

import (
	"bytes"
	"go_server/config"
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

func RequestPasswordReset(c *fiber.Ctx) error {
	var data = make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return utils.ErrorResponse(c, config.ErrorParsingBody, err)
	}

	email := data["email"]

	user, err := db.GetUserByEmail(email)

	if err != nil {
		return utils.ErrorResponse(c, config.UserNotFound, err)
	}

	EmailToken := &models.Email_Token{
		Email:     user.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		TokenType: "password_reset",
	}

	result, err := db.Create(config.EmailTokens, EmailToken)

	if err != nil {
		return utils.ErrorResponse(c, config.TokenCreatedFailed, err)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	EmailToken.ID = oid

	clientDomain := os.Getenv("CLIENT_DOMAIN")

	templateData := types.TemplateData{
		Domain: clientDomain,
		Token:  EmailToken.ID.Hex(),
	}

	templatePath := filepath.Join("templates", "password_reset_template.html")

	template, err := template.ParseFiles(templatePath)

	if err != nil {
		return utils.ErrorResponse(c, config.TemplateParseError, err)
	}

	var emailBody bytes.Buffer

	if err := template.Execute(&emailBody, templateData); err != nil {
		return utils.ErrorResponse(c, config.TemplateExecError, err)
	}

	if err := utils.SendEmail(user.Email, "Password Reset", emailBody.String()); err != nil {
		return utils.ErrorResponse(c, config.EmailSendError, err)
	}

	return utils.SuccessResponse(c, config.Success, nil)

}
