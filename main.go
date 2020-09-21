package main

import (
	"github.com/phongcao0114/book-api-gateway/resolver"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

const defaultPort = "9876"

func main() {
	schema := resolver.Init()

	http.Handle("/", playground.Handler("GraphQL playground", "/public"))
	http.Handle("/public", handler.New(&handler.Config{
		Schema: schema,
	}))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
