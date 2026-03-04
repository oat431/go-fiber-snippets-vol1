package bootstrap

import (
	"go-fiber-snippets/internal/delivery/http/controller"
	"go-fiber-snippets/internal/delivery/http/repository"
	"go-fiber-snippets/internal/delivery/http/service"
	"go-fiber-snippets/internal/graphql_api"

	"github.com/gofiber/fiber/v3/log"
)

type AppContainer struct {
	GraphQLController *controller.GraphQLController
}

func NewAppContainer() *AppContainer {

	bookRepo := repository.NewMockBookRepository()
	bookService := service.NewBookService(bookRepo)

	schema, err := graphql_api.SetupSchema(bookService)
	if err != nil {
		log.Fatal("Failed to setup GraphQL schema: ", err)
	}

	gqlController := controller.NewGraphQLController(schema)

	return &AppContainer{
		GraphQLController: gqlController,
	}
}
