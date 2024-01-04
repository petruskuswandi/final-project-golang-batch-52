package middleware

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/utils"
)

func UploadFile(ctx *gin.Context) {
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

	ctx.Set("filename", filename)

	ctx.Next()
}
