package graph

//go:generate go tool gqlgen generate

import "software-architecture/internal/core/ports"

// Resolver holds dependencies for GraphQL resolvers (driving adapter).
type Resolver struct {
	ProductService ports.ProductService
}
