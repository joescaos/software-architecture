package graphqladapter

import (
	"software-architecture/graph"
	"software-architecture/internal/core/ports"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
)

const graphqlPath = "/graphql"

// RegisterRoutes mounts GraphQL (POST) and the playground (GET) alongside existing REST routes.
func RegisterRoutes(router *gin.Engine, productService ports.ProductService) {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		ProductService: productService,
	}}))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})

	router.POST(graphqlPath, func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})
	router.GET(graphqlPath+"/playground", func(c *gin.Context) {
		playground.Handler("GraphQL playground", graphqlPath).ServeHTTP(c.Writer, c.Request)
	})
}
