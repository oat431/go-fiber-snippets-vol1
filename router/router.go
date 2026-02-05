package router

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func init() {
	log.Println("Creating router")
}

func StartServer() {
	app := fiber.New()
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get("/health", healthcheck.New())
	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
