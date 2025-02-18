//go:build tools

package tools

// TODO: use tools directive in go.mod
import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/air-verse/air"
	_ "github.com/google/wire/cmd/wire"
)
