//go:generate go run github.com/99designs/gqlgen v
package resolvers

import "github.com/shaheen-ct/go-hexagonal-architecture/internal/register"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Domain register.Registry
}
