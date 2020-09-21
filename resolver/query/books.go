package query

import (
	"encoding/json"
	"net/http"

	"github.com/phongcao0114/book-api-gateway/helper"

	"github.com/graphql-go/graphql"
	"github.com/phongcao0114/book-api-gateway/datatype"
	"github.com/phongcao0114/book-service/model"
)

var Books = &graphql.Field{
	Type:        datatype.ListBook,
	Description: "Books",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		url := "http://localhost:8080/api/v1/books"
		resp, err := helper.MakeRequest(http.MethodGet, url, nil)
		if err != nil {
			return []model.Book{}, err
		}
		defer resp.Body.Close()

		var books []model.Book
		json.NewDecoder(resp.Body).Decode(&books)
		return books, nil
	},
}
