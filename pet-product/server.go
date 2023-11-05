package main

import (
	"github.com/mhmmdFsl/my-online-petshop/pet-product/config"
	"github.com/mhmmdFsl/my-online-petshop/pet-product/service"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/mhmmdFsl/my-online-petshop/pet-product/graph"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	productService := service.NewProductService(&service.ProductServiceCfg{
		Collection: config.Collecion,
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ProductService: productService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
