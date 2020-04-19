package middlewares

// pulled straight from d docs

import (
"context"
"net/http"

// "github.com/vickywane/usecase-server/graph/db"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
    name string
}

// A stand-in for our database backed user object
type User struct {
    Name string
    IsAdmin bool
}

// Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware() func(http.Handler) http.Handler {
    // Database := db.Connect()

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            c, err := r.Cookie("auth-cookie")

            // Allow unauthenticated users in
            if err != nil || c == nil {
                next.ServeHTTP(w, r)
                return
            }

            // Todo IMPLEMENT  validateAndGetUserID & getUserById FIRST !!
            // userId, err := validateAndGetUserID(c)
            // if err != nil {
            //     http.Error(w, "Invalid cookie", http.StatusForbidden)
            //     return
            // }
            //
            // // get the user from the database
            // user := getUserByID(Database, userId)
            //
            // // put it in context
            // ctx := context.WithValue(r.Context(), userCtxKey, user)
            //
            // // and call the next with our new context
            // r = r.WithContext(ctx)
            // next.ServeHTTP(w, r)
        })
    }
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
    raw, _ := ctx.Value(userCtxKey).(*User)
    return raw
}


