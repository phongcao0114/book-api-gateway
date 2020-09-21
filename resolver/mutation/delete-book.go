package mutation

import (
	"net/http"

	"github.com/phongcao0114/book-api-gateway/helper"

	"github.com/graphql-go/graphql"
)

var DeleteBook = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "DeleteBook",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)
		url := "http://localhost:8080/api/v1/book/" + id
		resp, err := helper.MakeRequest(http.MethodDelete, url, nil)
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()
		return true, nil
	},
}
