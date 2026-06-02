package graphql

import (
	"go-fiber-snippets/internal/service"

	"github.com/graphql-go/graphql"
)

// SetupSchema wires all queries (and mutations) into a GraphQL schema.
func SetupSchema(bookSvc service.BookService) (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"books":  booksQuery(bookSvc),
			"book":   bookByIDQuery(bookSvc),

			// Add more queries here:
			// "users": usersQuery(userSvc),
		},
	})

	// Mutations (uncomment when ready):
	// rootMutation := graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "RootMutation",
	// 	Fields: graphql.Fields{
	// 		"createBook": createBookMutation(bookSvc),
	// 	},
	// })

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})
}
