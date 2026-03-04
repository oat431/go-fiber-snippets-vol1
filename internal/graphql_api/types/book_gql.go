package types

import "github.com/graphql-go/graphql"

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"title": &graphql.Field{Type: graphql.String},
	},
})
