package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func TestMiddleware(c *fiber.Ctx) error {
	log.Printf("Request on path: %s", c.Path())
	return c.Next() // Continue execution to the next middleware/route handler
}
