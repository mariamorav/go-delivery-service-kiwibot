package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/repositories"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/routes"
	"github.com/stretchr/testify/assert"
)

var router = routes.SetupRouter()
var botsRepo = repositories.NewBotRepository()

func TestCreateBot(t *testing.T) {

	type responseBody struct {
		Message string `json:"message"`
		Error   string `json:"error"`
		Bot     *models.Bot
	}

	testCases := []struct {
		bot        *models.Bot
		response   *responseBody
		StatusCode int
	}{
		{
			&models.Bot{
				Location: models.Location{
					Lat: 6.161398,
					Lon: -75.605353,
				},
				ZoneId: "zone-test",
			},
			&responseBody{
				Message: "Bot created successfully!",
				Bot: &models.Bot{
					Location: models.Location{
						Lat: 6.161398,
						Lon: -75.605353,
					},
					ZoneId: "zone-test",
				},
			},
			200,
		},
		{
			&models.Bot{
				Location: models.Location{
					Lat: 6.161398,
				},
				ZoneId: "zone-test",
			},
			&responseBody{
				Error: "Key: 'Bot.Location.Lon' Error:Field validation for 'Lon' failed on the 'required' tag",
			},
			500,
		},
		{
			&models.Bot{
				Location: models.Location{
					Lon: 6.161398,
				},
				ZoneId: "zone-test",
			},
			&responseBody{
				Error: "Key: 'Bot.Location.Lat' Error:Field validation for 'Lat' failed on the 'required' tag",
			},
			500,
		},
		{
			&models.Bot{
				Location: models.Location{
					Lat: 6.161398,
					Lon: 6.161398,
				},
			},
			&responseBody{
				Error: "Key: 'Bot.ZoneId' Error:Field validation for 'ZoneId' failed on the 'required' tag",
			},
			500,
		},
	}

	for _, testcase := range testCases {
		reqBody, _ := json.Marshal(testcase.bot)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/bots/new", bytes.NewReader(reqBody))
		router.ServeHTTP(w, req)

		var response *responseBody

		_ = json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, testcase.StatusCode, w.Code)
		assert.Contains(t, response.Error, testcase.response.Error)
		assert.Contains(t, response.Message, testcase.response.Message)

		if response.Bot != nil {
			_ = botsRepo.DeleteDoc(response.Bot.Id)
		}

	}

}

func TestGetBotsByZone(t *testing.T) {

	type responseBody struct {
		Total       int          `json:"total"`
		PreviousUrl string       `json:"previous_url"`
		NextUrl     string       `json:"next_url"`
		Results     []models.Bot `json:"results"`
	}

	zoneId := "zone-test"

	// Create a few bots for a zone
	bots := []*models.Bot{
		{
			Location: models.Location{
				Lat: 6.157655,
				Lon: -75.607373,
			},
			ZoneId: zoneId,
		},
		{
			Location: models.Location{
				Lat: 6.157655,
				Lon: -75.607373,
			},
			ZoneId: zoneId,
		},
	}

	var botsIds []string
	for _, bot := range bots {
		res, _ := botsRepo.Save(bot)
		botsIds = append(botsIds, res.Id)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/bots/"+zoneId+"?offset=0&limit=3", nil)
	router.ServeHTTP(w, req)

	var response *responseBody
	_ = json.Unmarshal(w.Body.Bytes(), &response)

	expectedNextURL := ""
	expectedPreviousURL := ""
	expectedTotal := 2

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, response.NextUrl, expectedNextURL)
	assert.Contains(t, response.PreviousUrl, expectedPreviousURL)
	assert.Equal(t, response.Total, expectedTotal)

	for _, id := range botsIds {
		_ = botsRepo.DeleteDoc(id)
	}

}

func TestAssignBotToOrder(t *testing.T) {

	type responseBody struct {
		Message string           `json:"message"`
		Bot     *models.Bot      `json:"bot"`
		Order   *models.Delivery `json:"order"`
	}

	// Create a Delivery
	delivery := models.Delivery{
		ZoneId: "us-central",
		Pickup: models.Pickup{
			PickupLat: 6.161375,
			PickupLon: -75.605641,
		},
		Dropoff: models.Dropoff{
			DropoffLat: 6.169699,
			DropoffLon: -75.591134,
		},
	}

	nearBot := models.Bot{
		Location: models.Location{
			Lat: 6.157655,
			Lon: -75.607373,
		},
		ZoneId: "us-central",
	}
	farBot := models.Bot{
		Location: models.Location{
			Lat: 6.152534,
			Lon: -75.613179,
		},
		ZoneId: "us-central",
	}

	deliveryRepo := repositories.NewDeliveryRepository()
	botsRepo := repositories.NewBotRepository()

	order, _ := deliveryRepo.Save(&delivery)
	// Create two bots one near and other one far from the pickup location
	botNearest, _ := botsRepo.Save(&nearBot)
	botFar, _ := botsRepo.Save(&farBot)

	orderId := models.BodyAssignBot{
		OrderId: order.Id,
	}

	reqBody, _ := json.Marshal(&orderId)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/bots/assign-order", bytes.NewReader(reqBody))
	router.ServeHTTP(w, req)

	var response *responseBody

	_ = json.Unmarshal(w.Body.Bytes(), &response)

	// Test must return that the assigned bot was the nearest.
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, response.Message, "order was assigned with success")
	assert.Contains(t, response.Order.State, "assigned")
	assert.Contains(t, response.Bot.Id, botNearest.Id)

	_ = botsRepo.DeleteDoc(botNearest.Id)
	_ = botsRepo.DeleteDoc(botFar.Id)
	_ = deliveryRepo.DeleteDoc(order.Id)

}
