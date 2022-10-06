package models

type Bot struct {
	Id       string   `json:"id"`
	Status   string   `json:"status"`
	Location Location `json:"location" validate:"required"`
	ZoneId   string   `json:"zone_id" validate:"required"`
}

type Location struct {
	Lat float64 `json:"lat" validate:"required"`
	Lon float64 `json:"lon" validate:"required"`
}
