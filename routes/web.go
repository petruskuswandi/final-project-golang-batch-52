package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/app/http/controllers"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/config"
)

func InitRoute(r *gin.Engine) {
	r.Group("")
	r.Static(config.STATIC_ROUTE, config.STATIC_DIR)

	// AUTH ROUTE
	r.POST("/login", controllers.Login)

	// ROUTE USER
	usersRoute := r.Group("users")
	userRoute := r.Group("user")
	usersRoute.GET("/paginate", controllers.GetUserPaginate)
	usersRoute.GET("", controllers.GetAllUser)
	usersRoute.POST("", controllers.Store)
	userRoute.GET("/:id", controllers.GetById)
	userRoute.PATCH("/:id", controllers.UpdateById)
	userRoute.DELETE("/:id", controllers.DeleteById)

	// ROUTE BOOK
	r.GET("/books", controllers.GetAllBook)

	// ROUTE FILE
	v1Route(r.Group(""))
}
