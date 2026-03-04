package queries

import (
	"go-fiber-snippets/internal/delivery/http/service"
	"go-fiber-snippets/internal/graphql_api/types"

	"github.com/graphql-go/graphql"
)

func GetBooksQuery(bookService service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.BookType),
		Description: "Get all books from database",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return bookService.GetAllBooks()
		},
	}
}

func GetBookByIDQuery(bookService service.BookService) *graphql.Field {
	return &graphql.Field{
		Type:        types.BookType,
		Description: "Get a single book by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			return bookService.GetBookByID(id)
		},
	}
}
