//go:generate go run github.com/99designs/gqlgen

package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/reearth/scaffold/server/internal/usecase"
	"github.com/vektah/gqlparser/v2/ast"
)

func NewServer(u usecase.Usecases, dev bool) *handler.Server {
	resolver := NewResolver(u)
	srv := handler.New(NewExecutableSchema(Config{Resolvers: resolver}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	if dev {
		srv.Use(extension.Introspection{})
	}

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}

func Playground(endpoint string) http.HandlerFunc {
	return playground.Handler("GraphQL playground", endpoint)
}
