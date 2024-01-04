package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/constanta"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/utils"
)

func SendStatus(ctx *gin.Context) {
	fileName := ctx.MustGet("filename").(string)

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "file uploaded",
		"file_name": fileName,
	})
}

func HandleUploadFile(ctx *gin.Context) {
	claimsData := ctx.MustGet("claimsData").(jwt.MapClaims)
	fmt.Println("claimsData => email => ", claimsData["email"])

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userId => ", userId)

	fileHeader, _ := ctx.FormFile("file")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file is required",
		})
		return
	}

	// validation file by extension
	fileExtention := []string{".png", ".jpg", ".jpeg", ".pdf"}

	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtention)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "file not allowed",
		})

		return
	}

	// validation file by content-type
	// fileType := []string{"image/jpg", "image/png", "image/jpeg", "image/svg"}

	// isFileValidated := utils.FileValidation(fileHeader, fileType)
	// if !isFileValidated {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"message": "file not allowed",
	// 	})

	// 	return
	// }

	extensionFile := filepath.Ext(fileHeader.Filename)

	filename := utils.RandomFileName(extensionFile)

	isSaved := utils.SaveFile(ctx, fileHeader, filename)
	if !isSaved {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error, can't save file",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file uploaded",
	})
}

func HandleRemoveFile(ctx *gin.Context) {
	filename := ctx.Param("filename")
	if filename == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file name is required",
		})
	}
	err := utils.RemoveFile(constanta.DIR_FILE + filename)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file deleted",
	})
}
