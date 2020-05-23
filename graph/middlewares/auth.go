package middlewares

import (
    "context"
    "errors"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
    "github.com/gin-gonic/gin"
    "github.com/go-pg/pg/v9"
    "github.com/joho/godotenv"
    "net/http"
    "strings"

    "github.com/vickywane/event-server/graph/model"
    "github.com/vickywane/event-server/graph/validators"
)

var Env, _ = godotenv.Read(".env")
var Key = Env["SECRET_KEY"]

type User struct {
    *pg.DB
}

func (U *User) GetUserByName(name string) (interface{}, error) {
    user := &model.User{}
    err := U.DB.Model(user).Where("name = ?", name).First()

    if err != nil {
        return nil, validators.NotFound
    }

    return user, err
}

func parseToken(Req *http.Request) (*jwt.Token, error) {
    jwtToken, err := request.ParseFromRequest(Req, extractor, func(token *jwt.Token) (interface{}, error) {
        jwtSecret := []byte(Env["SECRET_KEY"])
        return jwtSecret, nil
    })

    if err != nil {
        fmt.Println(err)
        return nil, validators.ParseToken
    }
    return jwtToken, nil
}

func JWT(U User) gin.HandlerFunc {
    return func(c *gin.Context) {
        Req := c.Request // contains http req details
        token, Terr := parseToken(Req)
        if Terr != nil {
            // http.Handler.ServeHTTP(c.Writer, c.Request)
            fmt.Println("an error with JWT middleware")
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            fmt.Println(token.Valid)
            fmt.Println(claims["jti"].(string), "valid token")
            return
        }

        user, _ := U.GetUserByName(claims["jti"].(string))

        fmt.Println(user, "\n")

        ctx := context.WithValue(Req.Context(), ContextKey, user)
        fmt.Println(ctx, "ctx value")

        // // gin error
        // http.Handler.ServeHTTP(c.Writer, Req.WithContext(ctx))
    }
}

var headerExtractor = &request.PostExtractionFilter{
    Extractor: request.HeaderExtractor{"Authorization"},
    Filter:    getBearFromTokenPayload,
}

func getBearFromTokenPayload(token string) (string, error) {
    bearer := "BEARER"
    if strings.ToUpper(token[0:len(bearer)]) == bearer {
        // trying to get the token after the BEARER keyword + a whitespace
        return token[len(bearer)+1:], nil
    }

    return token, nil
}

var extractor = &request.MultiExtractor{
    headerExtractor,
    request.ArgumentExtractor{"access_token"},
}

func ExtractCurrentUserFromContext(ctx context.Context) (*model.User, error) {
    User, ok := ctx.Value(ContextKey).(*model.User)
    if ctx.Value(ContextKey) == nil {
        return nil, errors.New("user Context Empty. No User here")
    }

    if !ok || User.ID == 0 {
        return nil, errors.New("user Context Empty. No User here")
    }

    return User, nil
}
