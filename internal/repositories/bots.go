package repositories

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/database"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"google.golang.org/api/iterator"
)

type BotRepository interface {
	Save(bot *models.Bot) (*models.Bot, error)
	FindByFilter(filterField, filterValue string, offset, limit int) ([]models.Bot, int, error)
	FindBotsAvailablesInZone(zoneId string) ([]*models.Bot, error)
	UpdateBotState(botId, newStatus string) (*models.Bot, error)
}

type botRepo struct{}

//NewBotRepository
func NewBotRepository() BotRepository {
	return &botRepo{}
}

const botsCollectionName = "bots"

func (*botRepo) Save(bot *models.Bot) (*models.Bot, error) {
	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	bot.Status = "available"

	docRef, _, err := client.Collection(botsCollectionName).Add(ctx, map[string]interface{}{
		"status": bot.Status,
		"location": map[string]float64{
			"lat": bot.Location.Lat,
			"lon": bot.Location.Lon,
		},
		"zone_id": bot.ZoneId,
	})
	if err != nil {
		return nil, err
	}

	bot.Id = docRef.ID

	return bot, nil

}

func (*botRepo) FindByFilter(filterField, filterValue string, offset, limit int) ([]models.Bot, int, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, 0, err
	}
	defer client.Close()

	var bots []models.Bot

	total, err := client.Collection(botsCollectionName).Where(filterField, "==", filterValue).Documents(ctx).GetAll()

	iter := client.Collection(botsCollectionName).Where(filterField, "==", filterValue).Offset(offset).Limit(limit).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		bot := models.Bot{
			Id:     doc.Ref.ID,
			Status: doc.Data()["status"].(string),
			Location: models.Location{
				Lat: doc.Data()["location"].(map[string]interface{})["lat"].(float64),
				Lon: doc.Data()["location"].(map[string]interface{})["lon"].(float64),
			},
			ZoneId: doc.Data()["zone_id"].(string),
		}

		bots = append(bots, bot)
	}

	return bots, len(total), nil

}

func (*botRepo) FindBotsAvailablesInZone(zoneId string) ([]*models.Bot, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	var bots []*models.Bot

	iter := client.Collection(botsCollectionName).Where("status", "==", "available").Where("zone_id", "==", zoneId).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			if len(bots) == 0 {
				return nil, fmt.Errorf("There are not available bots in the order zone")
			}
			break
		}

		bot := &models.Bot{
			Id:     doc.Ref.ID,
			Status: doc.Data()["status"].(string),
			Location: models.Location{
				Lat: doc.Data()["location"].(map[string]interface{})["lat"].(float64),
				Lon: doc.Data()["location"].(map[string]interface{})["lon"].(float64),
			},
			ZoneId: doc.Data()["zone_id"].(string),
		}

		bots = append(bots, bot)
	}

	return bots, nil

}

func (*botRepo) UpdateBotState(botId, newStatus string) (*models.Bot, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	_, err = client.Collection(botsCollectionName).Doc(botId).Update(ctx, []firestore.Update{
		{
			Path:  "status",
			Value: newStatus,
		},
	})

	if err != nil {
		return nil, err
	}

	dsnap, err := client.Collection(botsCollectionName).Doc(botId).Get(ctx)
	if err != nil {
		return nil, err
	}

	result := &models.Bot{
		Id:     dsnap.Ref.ID,
		Status: dsnap.Data()["status"].(string),
		Location: models.Location{
			Lat: dsnap.Data()["location"].(map[string]interface{})["lat"].(float64),
			Lon: dsnap.Data()["location"].(map[string]interface{})["lon"].(float64),
		},
		ZoneId: dsnap.Data()["zone_id"].(string),
	}

	return result, nil

}
