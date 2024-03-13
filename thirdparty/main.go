package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Gromitmugs/temporal-playground/thirdparty/internal/controller/resolvers"
	"github.com/Gromitmugs/temporal-playground/thirdparty/internal/graph/gqlgen"
	"github.com/go-chi/chi"
)

const defaultPort = "8001"

func main() {
	r := chi.NewRouter()
	srv := handler.NewDefaultServer(gqlgen.NewExecutableSchema(gqlgen.Config{
		Resolvers: &resolvers.Resolver{
			DB:          make(map[int]string),
			DBLastIndex: 0,
		},
	}))
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, r))
}
