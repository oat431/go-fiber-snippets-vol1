package router

import (
	"log"
	"os"

	"go-fiber-snippets/controller"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Creating router")
}

func StartServer() {
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/hello", controller.HelloWorld)

	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())

	// Additional routes can be added here

	env := godotenv.Load()
	if env != nil {
		log.Println("Error loading .env file")
	}
	port := ":" + os.Getenv("PORT")
	err := app.Listen(port)
	if err != nil {
		return
	}
}
