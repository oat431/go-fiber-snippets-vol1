package graphql

import (
	"github.com/graphql-go/graphql"
)

// BookType is the GraphQL representation of the Book domain model.
var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"title": &graphql.Field{Type: graphql.String},
	},
})
