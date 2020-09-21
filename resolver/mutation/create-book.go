package mutation

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/phongcao0114/book-api-gateway/helper"
	"github.com/phongcao0114/book-service/model"
)

var CreateBook = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "CreateBook",
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"author": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		author := p.Args["author"].(string)
		book := model.Book{
			Name:   name,
			Author: author,
		}

		jsonData, err := json.Marshal(&book)
		if err != nil {
			return false, err
		}

		url := "http://localhost:8080/api/v1/book"
		resp, err := helper.MakeRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()

		return true, nil
	},
}
