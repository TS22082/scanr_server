package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func ModifyRequestMiddleware(c *fiber.Ctx) error {
	// Read the request body into a map or struct

	var requestData map[string]interface{}
	if err := json.Unmarshal(c.Body(), &requestData); err != nil {
		return err // Handle error
	}

	// Modify the requestData as needed
	requestData["newField"] = "newValue"

	// Optionally, store the modified data in the Fiber context
	c.Locals("modifiedRequestData", requestData)

	return c.Next()
}
