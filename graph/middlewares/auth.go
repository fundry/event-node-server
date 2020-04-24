package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"time"


	"github.com/vickywane/event-server/graph/model"
)

var (
	Env, _ = godotenv.Read(".env")
)

var Key = Env["SECRET_KEY"]

func LoginHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(Key)
	c.JSON(200, gin.H{
		"userID":   claims[Key],
		"userName": user.(*model.User).Name,
	})
}

var AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "test zone",
	Key:         []byte("secret key"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: Key,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		// if v, ok := data.(*model.User); ok {
		// 	return jwt.MapClaims{
		// 		identityKey: v.Name,
		// 	}
		// }
		return jwt.MapClaims{}
	},

	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &model.User{
			Name: claims[Key].(string),
		}
	},

	Authenticator: func(c *gin.Context) (interface{}, error) {
		var UserDetails *model.User
		if err := c.ShouldBind(&UserDetails); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := "admin"
		password := "admin"

		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
			return &model.User{
				Name: userID,
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	},
})
