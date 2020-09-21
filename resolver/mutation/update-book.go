package mutation

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/phongcao0114/book-service/model"

	"github.com/phongcao0114/book-api-gateway/helper"
)

var UpdateBook = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Update book",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"author": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["id"].(string)
		name, _ := p.Args["name"].(string)
		author, _ := p.Args["author"].(string)
		book := model.Book{
			Name:   name,
			Author: author,
		}
		jsonData, err := json.Marshal(book)
		if err != nil {
			return false, err
		}
		url := "http://localhost:8080/api/v1/book/" + id
		resp, err := helper.MakeRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
		if err != nil {
			return false, err
		}
		defer resp.Body.Close()
		return true, nil
	},
}
