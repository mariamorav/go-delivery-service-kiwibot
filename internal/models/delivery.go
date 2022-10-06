package models

import "time"

type Delivery struct {
	Id           string    `json:"id" firestore:"id,omitempty"`
	CreationDate time.Time `json:"creation_date"`
	State        string    `json:"state"`
	Pickup       Pickup    `json:"pickup" validate:"required"`
	Dropoff      Dropoff   `json:"dropoff" validate:"required"`
	ZoneId       string    `json:"zone_id" validate:"required"`
}

type Pickup struct {
	PickupLat float64 `json:"pickup_lat" validate:"required"`
	PickupLon float64 `json:"pickup_lon" validate:"required"`
}

type Dropoff struct {
	DropoffLat float64 `json:"dropoff_lat" validate:"required"`
	DropoffLon float64 `json:"dropoff_lon" validate:"required"`
}
