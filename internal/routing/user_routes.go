// Package router provides the setup for the user routes in the web application.
// It utilizes the Fiber web framework to define and manage these routes.
package router

import (
	"go_server/internal/handlers/user_handler"
	"go_server/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetupUserRoutes configures the user-related routes for the web application.
// It creates a new route group prefixed with "/user" and applies a timeout
// middleware to all routes within this group. The timeout is set to 5 seconds.
// Currently, it defines a POST route for creating new users.
//
// Parameters:
//   - app: A *fiber.App instance which represents the Fiber application.
//
// Usage:
//
//	app := fiber.New()
//	API.SetupUserRoutes(app)
func SetupUserRoutes(app *fiber.App) {
	timeoutDuration := 5 * time.Second

	// Creating a route group for user-related endpoints.
	// Applying TimeoutMiddleware to handle request timeouts for all routes in this group.
	userGroup := app.Group("/api/user", middleware.TimeoutMiddleware(timeoutDuration))

	// POST endpoint for creating a new user.
	// The handler for this endpoint is defined in the user_handler package.
	userGroup.Post("/", user_handler.CreateNewUser)

	// POST endpoint for verifying a new user.
	// The handler for this endpoint is defined in the user_handler package.
	userGroup.Post("/verify/:token", user_handler.VerifyNewUser)
}
