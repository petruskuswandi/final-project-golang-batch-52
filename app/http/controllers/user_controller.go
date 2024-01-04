package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/requests"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/responses"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/models"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/database"
)

func GetAllUser(ctx *gin.Context) {
	users := new([]models.User)

	err := database.DB.Table("users").Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(responses.UserResponse)

	err := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data transmitted",
		"data":    user,
	})
}

func Store(ctx *gin.Context) {
	userReq := new(requests.UserRequest)

	err := ctx.ShouldBind(&userReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	userEmailExist := new(models.User)
	database.DB.Table("users").Where("email = ?", userReq.Email).First(&userEmailExist)

	if userEmailExist.Email != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "email already used",
		})

		return
	}

	user := new(models.User)
	user.Name = &userReq.Name
	user.Address = &userReq.Address
	user.Email = &userReq.Email
	user.BornDate = &userReq.BornDate

	errDb := database.DB.Table("users").Create(&user).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't create data",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data saved successfully.",
		"data":    user,
	})
}

func UpdateById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	userReq := new(requests.UserRequest)

	userEmailExist := new(models.User)

	errReq := ctx.ShouldBind(&userReq)

	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})

		return
	}

	errDb := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})

		return
	}

	// email exist
	errUserEmailExist := database.DB.Table("users").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email already used",
		})

		return
	}

	user.Name = &userReq.Name
	user.Email = &userReq.Email
	user.Address = &userReq.Address
	user.BornDate = &userReq.BornDate

	errUpdate := database.DB.Table("users").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "can't update data",
		})

		return
	}

	userResponse := responses.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data updated successfully",
		"data":    userResponse,
	})
}

func DeleteById(ctx *gin.Context) {
	id := ctx.Param("id")

	user := new(models.User)

	errFind := database.DB.Table("users").Where("id = ?", id).Find(&user).Error
	if errFind != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	if user.ID == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "data not found",
		})

		return
	}

	err := database.DB.Table("users").Unscoped().Where("id = ?", id).Delete(&user).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "data deleted successfully",
	})
}

func GetUserPaginate(ctx *gin.Context) {
	page := ctx.Query("page")
	if page == "" {
		page = "1"
	}

	perPage := ctx.Query("perPage")
	if perPage == "" {
		perPage = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	perPageInt, _ := strconv.Atoi(perPage)
	if pageInt < 1 {
		pageInt = 1
	}

	users := new([]models.User)

	err := database.DB.Table("users").Offset((pageInt - 1) * perPageInt).Limit(perPageInt).Find(&users).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     users,
		"page":     pageInt,
		"per_page": perPageInt,
	})
}
