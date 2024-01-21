package router

import (
	"go_server/internal/handlers"
	"go_server/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	timeoutDuration := 5 * time.Second

	authGroup := app.Group("/api/auth", middleware.TimeoutMiddleware(timeoutDuration))

	authGroup.Post("/emaillogin", handlers.EmailLogin)
}
