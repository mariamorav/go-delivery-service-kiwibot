package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/handlers"
)

func addDeliveriesRoutes(rg *gin.RouterGroup) {

	deliveries := rg.Group("/deliveries")

	deliveries.POST("/new", handlers.CreateDelivery)
	deliveries.GET("/:id", handlers.GetDeliveryById)
	deliveries.GET("/", handlers.GetDeliveries)
	deliveries.GET("/assing-pending-orders", handlers.AssignBotsToPendingOrders)

}
