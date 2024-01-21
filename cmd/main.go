package main

import (
	"go_server/internal/db"
	router "go_server/internal/routing"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	// Connect to MongoDB
	db.Connect()
	// Create a new Fiber instance
	app := fiber.New()
	// set up cors permissions
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://anotherdomain.com",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))
	// set up logging
	app.Use(logger.New())
	// Set up routes
	router.SetupRoutes(app)
	// Start the server
	err = app.Listen(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	// Graceful shutdown
	err = app.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
