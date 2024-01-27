package routing

import (
	"go_server/internal/handlers"
	"go_server/internal/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	timeoutDuration := 5 * time.Second

	auth := app.Group("/api/auth", middleware.TimeoutMiddleware(timeoutDuration))
	user := app.Group("/api/user", middleware.TimeoutMiddleware(timeoutDuration))

	auth.Post("/emaillogin", handlers.EmailLogin)

	user.Post("/", handlers.CreateNewUser)
	user.Post("/verify/:token", handlers.VerifyNewUser)
	user.Get("verifyJWT/:jwt_token", handlers.VerifyJWT)
	user.Get("/me", middleware.CheckJWT, handlers.GetMe)
}
