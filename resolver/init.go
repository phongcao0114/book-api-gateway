package resolver

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/phongcao0114/book-api-gateway/resolver/mutation"
	"github.com/phongcao0114/book-api-gateway/resolver/query"
)

func Init() *graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.Query(),
			Mutation: mutation.Mutation(),
		},
	)

	if err != nil {
		log.Fatalf("schema: Init: %v", err)
		return nil
	}

	return &schema
}
