package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/repositories"
	"github.com/stretchr/testify/assert"
)

var deliveryRepo = repositories.NewDeliveryRepository()
var resultIds []string

func TestCreateDelivery(t *testing.T) {
	type responseBody struct {
		Error    string           `json:"error"`
		Message  string           `json:"message"`
		Delivery *models.Delivery `json:"delivery"`
	}

	testCases := []struct {
		delivery   *models.Delivery
		response   *responseBody
		StatusCode int
	}{
		{
			&models.Delivery{
				ZoneId: "zone-test",
				Pickup: models.Pickup{
					PickupLat: 6.161375,
					PickupLon: -75.605641,
				},
				Dropoff: models.Dropoff{
					DropoffLat: 6.169699,
					DropoffLon: -75.591134,
				},
			},
			&responseBody{
				Message: "The order was created successfully",
				Delivery: &models.Delivery{
					ZoneId: "zone-test",
					Pickup: models.Pickup{
						PickupLat: 6.161375,
						PickupLon: -75.605641,
					},
					Dropoff: models.Dropoff{
						DropoffLat: 6.169699,
						DropoffLon: -75.591134,
					},
				},
			},
			200,
		},
		{
			&models.Delivery{
				ZoneId: "zone-test",
				Pickup: models.Pickup{
					PickupLat: 6.161375,
					PickupLon: -75.605641,
				},
				Dropoff: models.Dropoff{
					DropoffLat: 6.169699,
				},
			},
			&responseBody{
				Error: "Key: 'Delivery.Dropoff.DropoffLon' Error:Field validation for 'DropoffLon' failed on the 'required' tag",
			},
			500,
		},
	}

	for _, testcase := range testCases {
		reqBody, _ := json.Marshal(testcase.delivery)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/deliveries/new", bytes.NewReader(reqBody))
		router.ServeHTTP(w, req)

		var response *responseBody

		_ = json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, testcase.StatusCode, w.Code)
		assert.Contains(t, response.Error, testcase.response.Error)
		assert.Contains(t, response.Message, testcase.response.Message)

		if response.Delivery != nil {
			_ = deliveryRepo.DeleteDoc(response.Delivery.Id)
		}
	}

}

func TestGetDeliveryById(t *testing.T) {

	// create a delivery order
	delivery := models.Delivery{
		ZoneId: "zone-test",
		Pickup: models.Pickup{
			PickupLat: 6.161375,
			PickupLon: -75.605641,
		},
		Dropoff: models.Dropoff{
			DropoffLat: 6.169699,
			DropoffLon: -75.591134,
		},
	}

	result, _ := deliveryRepo.Save(&delivery)
	// Use endpoint to get it
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/deliveries/"+result.Id, nil)
	router.ServeHTTP(w, req)

	var response struct {
		Delivery *models.Delivery `json:"delivery"`
	}
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, response.Delivery.Id, result.Id)
	_ = deliveryRepo.DeleteDoc(delivery.Id)

}

func TestGetDeliveries(t *testing.T) {

	type responseBody struct {
		Total       int          `json:"total"`
		PreviousUrl string       `json:"previous_url"`
		NextUrl     string       `json:"next_url"`
		Results     []models.Bot `json:"results"`
	}

	// Create a few orders
	deliveries := []*models.Delivery{
		{
			ZoneId: "zone-test",
			Pickup: models.Pickup{
				PickupLat: 6.161375,
				PickupLon: -75.605641,
			},
			Dropoff: models.Dropoff{
				DropoffLat: 6.169699,
				DropoffLon: -75.591134,
			},
		},
		{
			ZoneId: "zone-test",
			Pickup: models.Pickup{
				PickupLat: 6.1614755,
				PickupLon: -74.6056451,
			},
			Dropoff: models.Dropoff{
				DropoffLat: 6.1696959,
				DropoffLon: -75.5911345,
			},
		},
	}

	for _, delivery := range deliveries {
		result, _ := deliveryRepo.Save(delivery)
		resultIds = append(resultIds, result.Id)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/deliveries/?offset=0&limit=3&order=desc", nil)
	router.ServeHTTP(w, req)

	var response *responseBody
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, response.Results[0].Id, resultIds[1])
	assert.Contains(t, response.Results[1].Id, resultIds[0])

}

func TestAssignBotsToPendingOrders(t *testing.T) {

	// Use the deliveries created in the previous test
	// Create bots in the zone
	bots := []*models.Bot{
		{
			Location: models.Location{
				Lat: 6.157655,
				Lon: -75.607373,
			},
			ZoneId: "zone-test",
		},
		{
			Location: models.Location{
				Lat: 6.134955,
				Lon: -75.603373,
			},
			ZoneId: "zone-test",
		},
	}

	var botsIds []string
	for _, bot := range bots {
		res, _ := botsRepo.Save(bot)
		botsIds = append(botsIds, res.Id)
	}

	// verify that bots were assigned succesfully
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/deliveries/assing-pending-orders", nil)
	router.ServeHTTP(w, req)

	var response *struct {
		Message string `json:"message"`
	}
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, response.Message, "all pending orders were assigned to the neareast bot available")

	// remove test docs
	for _, id := range resultIds {
		_ = deliveryRepo.DeleteDoc(id)
	}

	for _, id := range botsIds {
		_ = botsRepo.DeleteDoc(id)
	}

}
