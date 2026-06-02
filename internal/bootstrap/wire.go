package bootstrap

import (
	"go-fiber-snippets/internal/graphql"
	"go-fiber-snippets/internal/handler"
	"go-fiber-snippets/internal/repository"
	"go-fiber-snippets/internal/service"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// Container holds all wired dependencies.
type Container struct {
	GraphQLHandler fiber.Handler
	BookHandler    *handler.BookHandler
}

// NewContainer wires the full dependency graph.
func NewContainer() *Container {
	// --- Repository layer ---
	bookRepo := repository.NewMockBookRepo()

	// --- Service layer ---
	bookSvc := service.NewBookService(bookRepo)

	// --- HTTP handlers ---
	bookHandler := handler.NewBookHandler(bookSvc)

	// --- GraphQL ---
	schema, err := graphql.SetupSchema(bookSvc)
	if err != nil {
		log.Fatalf("graphql schema: %v", err)
	}

	return &Container{
		GraphQLHandler: graphql.Handler(schema),
		BookHandler:    bookHandler,
	}
}
