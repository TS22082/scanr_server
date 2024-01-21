package middleware

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

// TimeoutMiddleware creates a middleware for handling timeouts.
func TimeoutMiddleware(timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(c.Context(), timeout)
		defer cancel()

		// Create a channel to signal the completion of the handler
		done := make(chan struct{}, 1)

		// Run the handler in a goroutine
		go func() {
			defer close(done)
			// Forward the Fiber context with the new context
			c.SetUserContext(ctx)
			c.Next() // Call the next middleware/handler
		}()

		// Wait for either the handler to finish or the timeout
		select {
		case <-ctx.Done():
			// Timeout occurred
			return c.SendStatus(fiber.StatusRequestTimeout)
		case <-done:
			// Handler finished
			return nil
		}
	}
}
