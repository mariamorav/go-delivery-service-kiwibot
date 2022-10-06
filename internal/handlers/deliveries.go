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
	deliveriesRepo repositories.DeliveryRepository = repositories.NewDeliveryRepository()
)

func CreateDelivery(c *gin.Context) {

	var delivery models.Delivery

	if err := c.BindJSON(&delivery); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationError := validator.New().Struct(delivery)
	if validationError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationError.Error()})
		fmt.Println(validationError)
		return
	}

	result, err := deliveriesRepo.Save(&delivery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "The order was created successfully",
		"delivery": result,
	})

}

func GetDeliveryById(c *gin.Context) {
	deliveryId := c.Params.ByName("id")

	result, err := deliveriesRepo.FindDocumentById(deliveryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("delivery with id does not exists %v", deliveryId)})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"delivery": result,
	})

}

func GetDeliveries(c *gin.Context) {

	var queryParams models.QueryParams

	if c.ShouldBind(&queryParams) == nil {
		log.Println(queryParams.Offset)
		log.Println(queryParams.Limit)
		log.Println(queryParams.Order)
	}

	deliveries, total, err := deliveriesRepo.FindAll(int(queryParams.Offset), int(queryParams.Limit), queryParams.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if deliveries == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "there are no documents in deliveries collection"})
		fmt.Println(err)
		return
	}

	offsetString := strconv.Itoa(int(queryParams.Offset - queryParams.Offset))
	limitString := strconv.Itoa(int(queryParams.Limit))

	previousURL := "http://localhost:" + os.Getenv("PORT") + "/api/v1/deliveries?offset=" + offsetString + "&limit=" + limitString

	if queryParams.Offset == 0 {
		previousURL = ""
	}

	newOffset := queryParams.Offset + queryParams.Limit
	nextURL := "http://localhost:" + os.Getenv("PORT") + "/api/v1/deliveries?offset=" + strconv.Itoa(int(newOffset)) + "&limit=" + limitString

	if newOffset >= uint32(total) {
		nextURL = ""
	}

	c.JSON(http.StatusOK, gin.H{
		"total":        total,
		"previous_url": previousURL,
		"next_url":     nextURL,
		"results":      deliveries,
	})

}

func AssignBotsToPendingOrders(c *gin.Context) {

	ordersUnassigned, err := AssignBotsToAllPendingOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if len(ordersUnassigned) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":           "some orders could not be assigned due available of bots",
			"orders_unassigned": ordersUnassigned,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all pending orders were assigned to the neareast bot available",
	})

}
