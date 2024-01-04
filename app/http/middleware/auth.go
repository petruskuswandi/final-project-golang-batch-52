package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/utils"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})

		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})

		return
	}

	claimsData, err := utils.DecodeToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})

		return
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])
	ctx.Set("user_name", claimsData["name"])
	ctx.Set("user_email", claimsData["email"])

	ctx.Next()
}

func TokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("X-Token")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthenticated",
		})

		return
	}

	if token != "123" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "token not valid",
		})

		return
	}
	ctx.Next()
}
