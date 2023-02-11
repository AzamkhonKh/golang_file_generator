package main

import (
	"report-generator/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	userRepo := controllers.New()
	r.GET("/visits", userRepo.GetVisits)
	r.GET("/xlsx", userRepo.GenerateExcel)

	return r
}
