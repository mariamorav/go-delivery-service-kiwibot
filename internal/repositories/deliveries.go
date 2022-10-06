package repositories

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/database"
	"github.com/mariamorav/go-delivery-service-kiwibot/internal/models"
	"google.golang.org/api/iterator"
)

type DeliveryRepository interface {
	Save(delivery *models.Delivery) (*models.Delivery, error)
	FindAll(offset, limit int, order string) ([]models.Delivery, int, error)
	FindDocumentById(id string) (*models.Delivery, error)
	FindPendingOrders() ([]models.Delivery, error)
	UpdateDeliveryState(deliveryId, newState string) (*models.Delivery, error)
}

type deliveryRepo struct{}

//NewDeliveryRepository
func NewDeliveryRepository() DeliveryRepository {
	return &deliveryRepo{}
}

const (
	collectionName = "deliveries"
)

func (*deliveryRepo) Save(delivery *models.Delivery) (*models.Delivery, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	delivery.CreationDate = time.Now()
	delivery.State = "pending"

	docRef, result, err := client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"creation_date": delivery.CreationDate,
		"state":         delivery.State,
		"pickup": map[string]float64{
			"pickup_lat": delivery.Pickup.PickupLat,
			"pickup_lon": delivery.Pickup.PickupLon,
		},
		"dropoff": map[string]float64{
			"dropoff_lat": delivery.Dropoff.DropoffLat,
			"dropoff_lon": delivery.Dropoff.DropoffLon,
		},
		"zone_id": delivery.ZoneId,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Write Result: ", result)
	delivery.Id = docRef.ID

	return delivery, nil

}

func (*deliveryRepo) FindAll(offset, limit int, order string) ([]models.Delivery, int, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, 0, err
	}
	defer client.Close()

	var deliveries []models.Delivery

	total, err := client.Collection(collectionName).Documents(ctx).GetAll()

	queryOrder := firestore.Asc
	if order == "desc" {
		queryOrder = firestore.Desc
	}

	iter := client.Collection(collectionName).OrderBy("creation_date", queryOrder).Offset(offset).Limit(limit).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		//fmt.Println("Delivery data: ", doc.Data())

		delivery := models.Delivery{
			Id:           doc.Ref.ID,
			CreationDate: doc.Data()["creation_date"].(time.Time),
			State:        doc.Data()["state"].(string),
			Pickup: models.Pickup{
				PickupLat: doc.Data()["pickup"].(map[string]interface{})["pickup_lat"].(float64),
				PickupLon: doc.Data()["pickup"].(map[string]interface{})["pickup_lon"].(float64),
			},
			Dropoff: models.Dropoff{
				DropoffLat: doc.Data()["dropoff"].(map[string]interface{})["dropoff_lat"].(float64),
				DropoffLon: doc.Data()["dropoff"].(map[string]interface{})["dropoff_lon"].(float64),
			},
			ZoneId: doc.Data()["zone_id"].(string),
		}

		deliveries = append(deliveries, delivery)
	}

	return deliveries, len(total), nil

}

func (*deliveryRepo) FindDocumentById(id string) (*models.Delivery, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	dsnap, err := client.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Document data: %v", dsnap.Data())

	result := &models.Delivery{
		Id:           dsnap.Ref.ID,
		CreationDate: dsnap.Data()["creation_date"].(time.Time),
		State:        dsnap.Data()["state"].(string),
		Pickup: models.Pickup{
			PickupLat: dsnap.Data()["pickup"].(map[string]interface{})["pickup_lat"].(float64),
			PickupLon: dsnap.Data()["pickup"].(map[string]interface{})["pickup_lon"].(float64),
		},
		Dropoff: models.Dropoff{
			DropoffLat: dsnap.Data()["dropoff"].(map[string]interface{})["dropoff_lat"].(float64),
			DropoffLon: dsnap.Data()["dropoff"].(map[string]interface{})["dropoff_lon"].(float64),
		},
		ZoneId: dsnap.Data()["zone_id"].(string),
	}

	return result, nil

}

func (*deliveryRepo) FindPendingOrders() ([]models.Delivery, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	var deliveries []models.Delivery

	iter := client.Collection(collectionName).Where("state", "==", "pending").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		delivery := models.Delivery{
			Id:           doc.Ref.ID,
			CreationDate: doc.Data()["creation_date"].(time.Time),
			State:        doc.Data()["state"].(string),
			Pickup: models.Pickup{
				PickupLat: doc.Data()["pickup"].(map[string]interface{})["pickup_lat"].(float64),
				PickupLon: doc.Data()["pickup"].(map[string]interface{})["pickup_lon"].(float64),
			},
			Dropoff: models.Dropoff{
				DropoffLat: doc.Data()["dropoff"].(map[string]interface{})["dropoff_lat"].(float64),
				DropoffLon: doc.Data()["dropoff"].(map[string]interface{})["dropoff_lon"].(float64),
			},
			ZoneId: doc.Data()["zone_id"].(string),
		}

		deliveries = append(deliveries, delivery)
	}

	return deliveries, nil

}

func (*deliveryRepo) UpdateDeliveryState(deliveryId, newState string) (*models.Delivery, error) {

	ctx := context.Background()
	client, err := database.GetFirebaseClient()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	defer client.Close()

	_, err = client.Collection(collectionName).Doc(deliveryId).Update(ctx, []firestore.Update{
		{
			Path:  "state",
			Value: newState,
		},
	})

	if err != nil {
		return nil, err
	}

	dsnap, err := client.Collection(collectionName).Doc(deliveryId).Get(ctx)
	if err != nil {
		return nil, err
	}

	result := &models.Delivery{
		Id:           dsnap.Ref.ID,
		CreationDate: dsnap.Data()["creation_date"].(time.Time),
		State:        dsnap.Data()["state"].(string),
		Pickup: models.Pickup{
			PickupLat: dsnap.Data()["pickup"].(map[string]interface{})["pickup_lat"].(float64),
			PickupLon: dsnap.Data()["pickup"].(map[string]interface{})["pickup_lon"].(float64),
		},
		Dropoff: models.Dropoff{
			DropoffLat: dsnap.Data()["dropoff"].(map[string]interface{})["dropoff_lat"].(float64),
			DropoffLon: dsnap.Data()["dropoff"].(map[string]interface{})["dropoff_lon"].(float64),
		},
		ZoneId: dsnap.Data()["zone_id"].(string),
	}

	return result, nil

}
