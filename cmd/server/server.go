package main

import (
	"log"
	"net/http"
	"os"

	"github.com/shaheen-ct/go-hexagonal-architecture/db/pgx"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/graph/generated"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/graph/resolvers"
	"github.com/shaheen-ct/go-hexagonal-architecture/internal/register"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "8080"

//gqlgen is based on a Schema first approach

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	rg := register.Register(func() register.RegistryOptions {
		return register.RegistryOptions{
			PgxOption: pgx.Option{
				Host:     "localhost",
				User:     "root",
				Password: "123456",
				DBName:   "meetup",
				SSLMode:  "disable",
				Port:     "5432",
			},
		}
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{Domain: rg}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/gql"))
	http.Handle("/gql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
