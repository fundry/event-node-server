package main

import (
    "fmt"
    "github.com/gin-gonic/gin"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gin-contrib/cors"

    "github.com/vickywane/event-server/graph/db"
    "github.com/vickywane/event-server/graph/generated"
    InternalMiddleware "github.com/vickywane/event-server/graph/middlewares"
    Resolver "github.com/vickywane/event-server/graph/resolvers"
)

var Key = "id"

func graphqlHandler() gin.HandlerFunc {

    // Todo: Push logs into a log file
    Database := db.Connect()
    Database.AddQueryHook(db.Logs{})

    h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
        Resolvers: &Resolver.Resolver{
            DB: Database,
        }}))

    InternalMiddleware.DataLoaderMiddleware(Database, h)

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
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

    r.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
        AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
        AllowHeaders: []string{"content-type"},
    }),
        gin.Recovery(),
        InternalMiddleware.GinContextToContextMiddleware(),
    )

    // test routes for auth from git samples
    r.POST("/login", InternalMiddleware.AuthMiddleware.LoginHandler)
    r.GET("/auth", InternalMiddleware.AuthMiddleware.RefreshHandler)
    r.GET("/hello", InternalMiddleware.LoginHandler)
    // ====================>

    r.POST("/query",
        graphqlHandler(),
    )
    fmt.Println("Playground running at http://localhost:4040")

    r.GET("/", playgroundHandler())
    r.Run(":4040")
}
