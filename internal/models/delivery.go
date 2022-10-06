package models

import "time"

type Delivery struct {
	Id           string    `json:"id" firestore:"id,omitempty"`
	CreationDate time.Time `json:"creation_date"`
	State        string    `json:"state"`
	Pickup       Pickup    `json:"pickup"`
	Dropoff      Dropoff   `json:"dropoff"`
	ZoneId       string    `json:"zone_id"`
}

type Pickup struct {
	PickupLat float64 `json:"pickup_lat"`
	PickupLon float64 `json:"pickup_lon"`
}

type Dropoff struct {
	DropoffLat float64 `json:"dropoff_lat"`
	DropoffLon float64 `json:"dropoff_lon"`
}
