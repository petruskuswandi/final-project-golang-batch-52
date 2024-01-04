package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/config"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/database"
	"gitlab.com/petruskuswandi1/final-project-golang-batch-52/routes"
)

func Boot() {
	// LOAD .env FILE
	godotenv.Load()

	// INIT CONFIG
	config.InitConfig()

	// DATABASE CONNECTION
	database.ConnectDatabase()

	config.DefaultLogging("logs/file/app.log")

	// INIT GIN ENGINE
	r := gin.Default()

	// CORS
	r.Use(config.CorsConfigContrib())
	// r.Use(config.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(r)

	// RUN APP
	port := fmt.Sprintf(":%s", config.PORT)
	if config.PORT == "" {
		r.Run()
	}
	r.Run(port)
}
