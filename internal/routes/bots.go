package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/handlers"
)

func addBotsRoutes(rg *gin.RouterGroup) {

	bots := rg.Group("/bots")

	bots.POST("/new", handlers.CreateBot)
	bots.GET("/:zone_id", handlers.GetBotsByZone)
	bots.POST("/assign-order", handlers.AssignBotToOrder)

}
