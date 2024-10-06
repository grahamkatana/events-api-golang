package middleware

import (
	"grahamkatana/api/events/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckIsTokenValid(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "error, token is missing in headers",
		})
		return
	}
	id, err := utils.VerifyJwtToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	ctx.Set("userId", id)
	ctx.Next()

}
