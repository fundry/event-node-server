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
	// Todo Find how to use Cors with GIN
	// h.Use(cors.New(cors.Options{
	// 	AllowedOrigins:         []string{"http://localhost:4040"},
	// 	AllowCredentials:       true,
	// 	Debug:                  true,
	// }))

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
		AllowOrigins:     []string{"https://http://localhost:8080"},
		AllowMethods:     []string{"GET", "PUT", "POST" , "DELETE"},
	}))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
