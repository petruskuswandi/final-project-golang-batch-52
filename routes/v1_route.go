package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/controllers"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/middleware"
)

func v1Route(app *gin.RouterGroup) {
	r := app

	authRoute := r.Group("file", middleware.AuthMiddleware)
	authRoute.POST("", controllers.HandleUploadFile)
	authRoute.POST("/middleware", middleware.UploadFile, controllers.SendStatus)
	authRoute.DELETE("/:filename", controllers.HandleRemoveFile)
}
