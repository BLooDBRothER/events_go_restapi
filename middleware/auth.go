package middleware

import (
	"net/http"

	"example.com/events/jwt"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if(token == "") {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User UnAuthorized"})
		return
	}

	userId, err := jwt.VerifyToken(token)

	if(err != nil) {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "User UnAuthorized"})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
