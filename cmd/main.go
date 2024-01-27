package main

import (
	"go_server/internal/db"
	"go_server/internal/routing"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	db.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173, https://anotherdomain.com",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(logger.New())
	routing.SetupRoutes(app)

	err = app.Listen(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	err = app.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
