package middlewares

import (
	"example.com/restful-api-example/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "request does not contain an access token"})
			ctx.Abort()
			return
		}
		user_name, err := auth.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("uun", user_name)

		ctx.Next()
	}
}
