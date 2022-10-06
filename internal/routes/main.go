package routes

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mariamorav/go-delivery-service-kiwibot/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.New()

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(gin.Logger())
	router.Use(cors.Default())
	getRoutes()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
