package utils

import (
	"fmt"
	constants "go_server/config"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// SendEmail sends an email to the specified email address.
func SendEmail(email string, subject string, body string) error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Set up email
	m := gomail.NewMessage()
	m.SetHeader("From", constants.DoNotReplyEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Set up dialer
	d := gomail.NewDialer("smtp.gmail.com", 587, constants.DoNotReplyEmail, os.Getenv("GOOGLE_APP_PW"))

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
