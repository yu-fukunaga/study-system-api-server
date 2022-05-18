package main

import (
	"log"
	"net/http"
	"os"
	"study-system/graph"
	"study-system/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"
const defaultHost = "localhost"

func startHttpServer() {
	port := os.Getenv("SAMPLE_PORT")
	if port == "" {
		port = defaultPort
	}
	host := os.Getenv("SAMPLE_HOST")
	if host == "" {
		port = defaultHost
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
