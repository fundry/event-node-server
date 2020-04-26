package middlewares

import (
    "context"
    "fmt"

    "github.com/gin-gonic/gin"
)

var ContextKey = "GinContextKey"
func GinContextToContextMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx := context.WithValue(c.Request.Context(), ContextKey, c)
        c.Request = c.Request.WithContext(ctx)
        c.Next()
    }
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
    ginContext := ctx.Value(ContextKey)
    if ginContext == nil {
        err := fmt.Errorf("could not retrieve gin.Context")
        return nil, err
    }

    gc, ok := ginContext.(*gin.Context)
    if !ok {
        err := fmt.Errorf("gin.Context has wrong type")
        return nil, err
    }
    return gc, nil
}