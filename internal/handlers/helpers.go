package handlers

import (
	"fmt"

	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"github.com/umahmood/haversine"
)

type QueryParams struct {
	Offset uint32 `form:"offset"`
	Limit  uint32 `form:"limit"`
	Order  string `form:"order"`
}

func GetNearestBotAvailable(delivery *models.Delivery) (*models.Bot, error) {

	bots, err := botsRepo.FindBotsAvailablesInZone(delivery.ZoneId)
	if err != nil {
		return nil, err
	}

	orderLat := delivery.Pickup.PickupLat
	orderLon := delivery.Pickup.PickupLon

	nearestBot := bots[0]

	pickupLocation := haversine.Coord{Lat: orderLat, Lon: orderLon}
	currentLocation := haversine.Coord{Lat: bots[0].Location.Lat, Lon: bots[0].Location.Lon}
	_, distanceinKm := haversine.Distance(pickupLocation, currentLocation)

	// @TODO: This operation could be done with go routines to improve performance
	for _, bot := range bots {
		botLat := bot.Location.Lat
		botLon := bot.Location.Lon

		pickupLocation := haversine.Coord{Lat: orderLat, Lon: orderLon}
		currentLocation := haversine.Coord{Lat: botLat, Lon: botLon}
		_, km := haversine.Distance(pickupLocation, currentLocation)
		//fmt.Println("Kilometers:", km)
		if km < distanceinKm {
			nearestBot = bot
			distanceinKm = km
		}
	}

	return nearestBot, nil

}

// @TODO: this process could be running on background searching availability of bots for pending orders.
func AssignBotsToAllPendingOrders() ([]*models.Delivery, error) {

	// Get pending orders
	pendingOrders, err := repo.FindPendingOrders()
	if err != nil {
		errMsg := fmt.Errorf("there are no pending orders: %v", err)
		return nil, errMsg
	}

	var ordererUnfilled []*models.Delivery

	for _, order := range pendingOrders {
		nearestBot, err := GetNearestBotAvailable(&order)
		if err != nil {
			fmt.Printf("there are no nearest bots available for the order %v", err)
			ordererUnfilled = append(ordererUnfilled, &order)
			continue
		}
		_, err = botsRepo.UpdateBotState(nearestBot.Id, "reserved")
		if err != nil {
			return nil, err
		}
		_, err = repo.UpdateDeliveryState(order.Id, "assigned")
		if err != nil {
			return nil, err
		}
	}

	return ordererUnfilled, err

}
