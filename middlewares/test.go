package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthorizeJWT is ...
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authHeader := ctx.GetHeader("Mock-auth")

		if authHeader == "benar" {
			fmt.Println("Authorized")
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token!",
			})
		}

	}
}
