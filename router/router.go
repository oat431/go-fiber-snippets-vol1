package router

import (
	"go-fiber-snippets/bootstap"
	"go-fiber-snippets/middleware"
	"log"
	"os"

	"go-fiber-snippets/controller"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/joho/godotenv"
)

func init() {
	log.Println("Creating router")
}

func StartServer() {
	scheduler, batchFailed := bootstap.RegisterJobs()
	if batchFailed != nil {
		log.Fatal("Failed to start scheduler: ", batchFailed)
	}
	scheduler.Start()

	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")
	redirect := v1.Group("/redirect")

	app.Use("/chat", middleware.WebSocketMiddleware)

	v1.Get("/hello", controller.HelloWorld)
	redirect.Get("/linkedin", controller.ToLinkedIn)
	redirect.Get("/github", controller.ToGitHub)
	redirect.Get("/facebook", controller.ToFacebook)

	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())

	app.Get("/chat", websocket.New(controller.WebSocketExample))

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
