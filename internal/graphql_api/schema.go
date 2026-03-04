package graphql_api

import (
	"go-fiber-snippets/internal/delivery/http/service"
	"go-fiber-snippets/internal/graphql_api/queries"

	"github.com/graphql-go/graphql"
)

func SetupSchema(bookService service.BookService) (graphql.Schema, error) {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"GetAllBook":  queries.GetBooksQuery(bookService),
			"GetBookById": queries.GetBookByIDQuery(bookService),

			// ตัวอย่างในอนาคต:
			// "users": queries.GetUsersQuery(userService),
			// "links": queries.GetLinksQuery(linkService),
		},
	})

	/* // ถ้ามี Mutations (Insert/Update) ก็ทำแบบเดียวกัน
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createBook": mutations.CreateBookMutation(bookService),
		},
	})
	*/

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
		// Mutation: rootMutation,
	})
}
