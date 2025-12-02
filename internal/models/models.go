package models

import "time"

type GeoLocation struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
}

type WeatherData struct {
	Name        string    `json:"name"`
	Temperature float64   `json:"temperature"`
	TimeStamp   time.Time `json:"time_stamp"`
}
