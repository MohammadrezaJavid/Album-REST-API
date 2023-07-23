package middlewares

import (
	"album/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.IndentedJSON(http.StatusUnauthorized,
				gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := authentication.ValideToken(tokenString)
		if err != nil {
			context.IndentedJSON(http.StatusUnauthorized,
				gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}

}
