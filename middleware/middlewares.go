package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authmiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, err := ctx.Request.BasicAuth()

		if username == "r" && password == "1" {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		}
	}
}
