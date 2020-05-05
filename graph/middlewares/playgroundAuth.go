package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func PlaygroundAuth() gin.HandlerFunc {
	godotenv.Load(".env")
	Envs, _ := godotenv.Read(".env")

	return gin.BasicAuth(gin.Accounts{
		Envs["PLAYGROUND_USER"]: Envs["PLAYGROUND_PASSWORD"],
	})
}
