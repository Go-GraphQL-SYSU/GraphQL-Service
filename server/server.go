package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	GraphQL_Service "github.com/Go-GraphQL-SYSU/GraphQL-Service"
	"github.com/Go-GraphQL-SYSU/GraphQL-Service/service"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/login", service.LoginHandler)
	http.HandleFunc("/logout", service.LogoutHandler)
	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(GraphQL_Service.NewExecutableSchema(GraphQL_Service.Config{Resolvers: &GraphQL_Service.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
