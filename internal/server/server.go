package server

import (
	"go-fiber-snippets/internal/handler"
	"go-fiber-snippets/internal/middleware"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

// Config holds everything needed to build the Fiber app.
type Config struct {
	Port           string
	GraphQLHandler fiber.Handler
	BookHandler    *handler.BookHandler
}

// New builds and returns a fully configured Fiber app.
func New(cfg Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "go-fiber-snippets",
		ErrorHandler: defaultErrorHandler,
	})

	// --- Global middleware ---
	app.Use(recover.New())
	app.Use(middleware.RequestLogger())

	// --- Health checks ---
	app.Get(healthcheck.LivenessEndpoint, healthcheck.New())
	app.Get(healthcheck.ReadinessEndpoint, healthcheck.New())
	app.Get(healthcheck.StartupEndpoint, healthcheck.New())

	// --- GraphQL ---
	if cfg.GraphQLHandler != nil {
		app.Post("/graphql", cfg.GraphQLHandler)
	}

	// --- Routes ---
	registerRoutes(app, cfg)

	return app
}

func registerRoutes(app *fiber.App, cfg Config) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Basic
	v1.Get("/hello", handler.Hello)

	// Books (REST: list + create)
	if cfg.BookHandler != nil {
		books := v1.Group("/books")
		books.Get("/", cfg.BookHandler.ListBooks)
		books.Post("/", cfg.BookHandler.CreateBook)
	}

	// Redirects
	redirect := v1.Group("/redirect")
	redirect.Get("/linkedin", handler.RedirectLinkedIn)
	redirect.Get("/github", handler.RedirectGitHub)
	redirect.Get("/facebook", handler.RedirectFacebook)

	// Webhook
	v1.Post("/hook", handler.Webhook)

	// WebSocket
	app.Use("/chat", middleware.WSGuard)
	app.Get("/chat", websocket.New(handler.WebSocketChat))
}

func defaultErrorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
