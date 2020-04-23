package main

import (
	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"

	"github.com/vickywane/event-server/graph/db"
	"github.com/vickywane/event-server/graph/generated"
	Resolver "github.com/vickywane/event-server/graph/resolvers"
	// InternalMiddleware "github.com/vickywane/event-server/graph/middlewares"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	Database := db.Connect()
	Database.AddQueryHook(db.Logs{})

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver.Resolver{
			DB: Database,
		}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders: []string{"content-type"},
	}))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
