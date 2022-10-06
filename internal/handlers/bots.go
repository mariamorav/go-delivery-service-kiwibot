package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/repositories"
)

var (
	botsRepo repositories.BotRepository = repositories.NewBotRepository()
)

// CreateBot godoc
// @Summary create a new bot
// @Schemes
// @Description This endpoint allows you to create a new bot with available status as default.
// @Tags bots
// @Accept json
// @Produce json
// @Param Body body models.CreateBotRequest true "The body request to create a Bot"
// @Success 200 {object} models.CreateBotResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /bots/new [post]
func CreateBot(c *gin.Context) {

	var bot models.Bot
	if err := c.BindJSON(&bot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationError := validator.New().Struct(bot)
	if validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationError.Error()})
		fmt.Println(validationError)
		return
	}

	result, err := botsRepo.Save(&bot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Bot created successfully!",
		"bot":     result,
	})

}

// GetBotsByZone godoc
// @Summary get bots by a zone_id
// @Schemes
// @Description This endpoint allows you to list the bots located in the zone_id submitted.
// @Tags bots
// @Accept json
// @Produce json
// @Param zone_id path string true "Id of the zone where you want to find bots"
// @Param offset query int true "Start point of documents search"
// @Param limit query int true	"Limit number of results returned by search"
// @Success 200 {object} models.GetBotsByZoneResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure 404 {object} models.MessageNotFound "there are no results for your search"
// @Router /bots/{zone_id} [get]
func GetBotsByZone(c *gin.Context) {

	var queryParams models.QueryParams

	if c.ShouldBind(&queryParams) == nil {
		log.Println(queryParams.Offset)
		log.Println(queryParams.Limit)
	}

	zoneId := c.Params.ByName("zone_id")

	results, total, err := botsRepo.FindByFilter("zone_id", zoneId, int(queryParams.Offset), int(queryParams.Limit))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if results == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "there are no results for your search"})
		fmt.Println(err)
		return
	}

	offsetString := strconv.Itoa(int(queryParams.Offset - queryParams.Offset))
	limitString := strconv.Itoa(int(queryParams.Limit))

	previousURL := "http://localhost:" + os.Getenv("PORT") + "/api/v1/bots/" + zoneId + "?offset=" + offsetString + "&limit=" + limitString

	if queryParams.Offset == 0 {
		previousURL = ""
	}

	newOffset := queryParams.Offset + queryParams.Limit
	nextURL := "http://localhost:" + os.Getenv("PORT") + "/api/v1/bots/" + zoneId + "?offset=" + strconv.Itoa(int(newOffset)) + "&limit=" + limitString

	if newOffset >= uint32(total) {
		nextURL = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"total":        total,
		"previous_url": previousURL,
		"next_url":     nextURL,
		"results":      results,
	})

}

// GetBotsByZone godoc
// @Summary assign a bot to an order
// @Schemes
// @Description This endpoint allows you assign an available bot to an a pending order.
// @Tags bots
// @Accept json
// @Produce json
// @Param Body body models.BodyAssignBot true "The body request to assign a bot"
// @Success 200 {object} models.AssignBotToOrderResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /bots/assign-order [post]
func AssignBotToOrder(c *gin.Context) {

	var req models.BodyAssignBot
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationError := validator.New().Struct(req)
	if validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationError.Error()})
		fmt.Println(validationError)
		return
	}

	// Assign available and nearest bot to the order
	// Get Order by Id
	deliveryOrder, err := deliveriesRepo.FindDocumentById(req.OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Get the nearest and available bot
	nearestBot, err := GetNearestBotAvailable(deliveryOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	// update bot state and order state
	updatedBot, err := botsRepo.UpdateBotState(nearestBot.Id, "reserved")
	updatedOrder, err := deliveriesRepo.UpdateDeliveryState(deliveryOrder.Id, "assigned")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "order was assigned with success",
		"bot":     updatedBot,
		"order":   updatedOrder,
	})

}
