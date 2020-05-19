package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/vickywane/event-server/graph/db"
	"github.com/vickywane/event-server/graph/generated"
	InternalMiddleware "github.com/vickywane/event-server/graph/middlewares"
	Resolver "github.com/vickywane/event-server/graph/resolvers"
)

var Key = "id"
var Database = db.Connect()

// Todo Decompress this file later!
func graphqlHandler() gin.HandlerFunc {
	// Todo: Push logs into a log file
	Database.AddQueryHook(db.Logs{})

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver.Resolver{
			DB: Database,
		}}))

	InternalMiddleware.DataLoaderMiddleware(Database, graphqlHandler)

	return func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	r := gin.Default()
	r.Use(
		cors.New(cors.Config{ // Todo am facing CORS issues now!
			// AllowOrigins: []string{"http://localhost:3000/", "http://localhost:4040",
			//     "http://localhost:8080"},
			AllowMethods:    []string{"GET", "PUT", "POST", "DELETE"},
			AllowHeaders:    []string{"content-type"},
			AllowAllOrigins: true,
		}),
		gin.Recovery(),
		InternalMiddleware.GinContextToContextMiddleware(),
		// this is restricting other services from acessing my endpoint
		//  InternalMiddleware.PlaygroundAuth(),
	)

	r.POST("/query",
		InternalMiddleware.JWT(InternalMiddleware.User{Database}),
		graphqlHandler(),
	)
	fmt.Println("Playground running at http://localhost:4040")

	r.GET("/", playgroundHandler())
	r.Run(":4040")
}
