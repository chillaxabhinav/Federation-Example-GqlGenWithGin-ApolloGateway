package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"product/graph/generated"
	graph "product/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"

	middleware "github.com/s12i/gin-throttle"
)

var runningPort string = ":4002"

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	return func(c *gin.Context) {

		cgql := generated.Config{Resolvers: &graph.Resolver{}}

		srv := handler.NewDefaultServer(generated.NewExecutableSchema(cgql))

		srv.Use(extension.FixedComplexityLimit(200))

		srv.ServeHTTP(c.Writer, c.Request)
	}

}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default() // By default runs in debug mode. Switch to "release" mode in production

	maxEventsPerSec := 5000

	maxBurstSize := 200

	router.Use(middleware.Throttle(maxEventsPerSec, maxBurstSize))

	router.POST("/query", graphqlHandler())

	router.GET("/product/playground", playgroundHandler())

	log.Println("Product Server running")

	router.Run(runningPort)
}
