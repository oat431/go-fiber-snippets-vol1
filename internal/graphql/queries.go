package graphql

import (
	"go-fiber-snippets/internal/service"

	"github.com/graphql-go/graphql"
)

// booksQuery returns all books.
//
// GraphQL usage:
//
//	{ books { id title } }
func booksQuery(bookSvc service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(BookType),
		Description: "Returns all books",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return bookSvc.GetAllBooks()
		},
	}
}

// bookByIDQuery returns a single book by ID.
//
// GraphQL usage:
//
//	{ book(id: "1") { id title } }
func bookByIDQuery(bookSvc service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        BookType,
		Description: "Returns a book by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, _ := p.Args["id"].(string)
			return bookSvc.GetBookByID(id)
		},
	}
}
