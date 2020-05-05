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
)

var Env, _ = godotenv.Read(".env")

var Key = Env["SECRET_KEY"]

type Resolver struct {
	DB *pg.DB
}

// used by auth middleware to get current user from ctx
func (r *Resolver) GetAUserById(id string) (*model.User, error) {
	user := model.User{}
	err := r.DB.Model(&user).Where("id = ?", id).First()
	return &user, err
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		Req := c.Request
		token, Terr := parseToken(Req)
		if Terr != nil {
			// http.Handler.ServeHTTP(c.Request, c.Writer)
			fmt.Println("an error with JWT middleware")
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if ok || !token.Valid {
			return
		}
		user, err := "", ""
		if err != "" {
			return
		}

		ctx := context.WithValue(Req.Context(), ContextKey, user)
		fmt.Println(ctx, "look ctx")
		// gin error
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

func parseToken(Req *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(Req, extractor, func(token *jwt.Token) (interface{}, error) {
		jwtSecret := Env["SECRET_KEY"]
		return jwtSecret, nil
	})

	if err != nil {
		return nil, errors.New("shit happened")
	}

	return jwtToken, nil
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
