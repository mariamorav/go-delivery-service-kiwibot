package models

type Bot struct {
	Id       string   `json:"id"`
	Status   string   `json:"status"`
	Location Location `json:"location"`
	ZoneId   string   `json:"zone_id"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
