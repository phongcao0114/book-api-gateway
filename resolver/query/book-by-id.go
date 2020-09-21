package query

import (
	"encoding/json"
	"net/http"

	"github.com/phongcao0114/book-api-gateway/helper"

	"github.com/graphql-go/graphql"
	"github.com/phongcao0114/book-api-gateway/datatype"
	"github.com/phongcao0114/book-service/model"
)

var BookByID = &graphql.Field{
	Type:        datatype.Book,
	Description: "BookByID",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)
		url := "http://localhost:8080/api/v1/book/" + id
		resp, err := helper.MakeRequest(http.MethodGet, url, nil)
		if err != nil {
			return model.Book{}, err
		}
		defer resp.Body.Close()

		var book model.Book
		json.NewDecoder(resp.Body).Decode(&book)
		return book, nil
	},
}
