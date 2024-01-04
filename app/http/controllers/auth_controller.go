package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/requests"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/models"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/database"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/utils"
)

func Login(ctx *gin.Context) {
	loginReq := new(requests.LoginRequest)
	err := ctx.ShouldBind(&loginReq)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user := new(models.User)
	errUser := database.DB.Table("users").Where("email = ?", loginReq.Email).Find(&user).Error
	if errUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid",
		})

		return
	}

	// check password
	if loginReq.Password != "12345" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "credential not valid",
		})

		return
	}

	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed generate token",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successfully",
		"token":   token,
	})
}
