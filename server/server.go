package main

import (
	GraphQL_Service "Go-GraphQL-SYSU/GraphQL-Service"
	"github.com/99designs/gqlgen/handler"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	http.HandleFunc("/login", GraphQL_Service.LoginHandler)
	http.HandleFunc("/logout", GraphQL_Service.LogoutHandler)
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(GraphQL_Service.NewExecutableSchema(GraphQL_Service.Config{Resolvers: &GraphQL_Service.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
