package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HasToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		if _, ok := context.GetQuery("token"); !ok {
			context.String(http.StatusUnauthorized, "Token is missing...")
		}
	}
}
