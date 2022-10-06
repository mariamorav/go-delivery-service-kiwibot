package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func Run() {
	getRoutes()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.Run(":" + port)
}

func getRoutes() {
	v1 := router.Group("/api/v1")
	addDeliveriesRoutes(v1)
	addBotsRoutes(v1)
}

func SetupRouter() *gin.Engine {
	getRoutes()
	return router
}
