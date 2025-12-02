package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/alonsoF100/weather-service/internal/models"
)

type Client struct {
	Client *http.Client
}

func New(cfg config.ClientConfig) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

func (c *Client) GetWeather(geoLocation *models.GeoLocation) (*models.WeatherData, error) {
	baseURL := "https://api.open-meteo.com/v1/forecast"

	params := url.Values{}
	params.Add("latitude", fmt.Sprintf("%f", geoLocation.Latitude))
	params.Add("longitude", fmt.Sprintf("%f", geoLocation.Longitude))
	params.Add("current", "temperature_2m")

	if geoLocation.Timezone != "" {
		params.Add("timezone", geoLocation.Timezone)
	} else {
		params.Add("timezone", "auto")
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("!StatusOK")
	}

	var temp struct {
		Current struct {
			Time          string  `json:"time"`
			Temperature2m float64 `json:"temperature_2m"`
		} `json:"current"`
	}

	if err := json.NewDecoder(response.Body).Decode(&temp); err != nil {
		return nil, err
	}

	var location *time.Location
	if geoLocation.Timezone != "" {
		loc, err := time.LoadLocation(geoLocation.Timezone)
		if err == nil {
			location = loc
		} else {
			location = time.UTC
		}
	} else {
		location = time.UTC
	}

	var timestamp time.Time
	if location != nil {
		timestamp, err = time.ParseInLocation("2006-01-02T15:04", temp.Current.Time, location)
		if err != nil {
			return nil, err
		}
	} else {
		timestamp, err = time.Parse("2006-01-02T15:04", temp.Current.Time)
		if err != nil {
			return nil, err
		}
	}

	return &models.WeatherData{
		Name:        geoLocation.Name,
		Temperature: temp.Current.Temperature2m,
		TimeStamp:   timestamp,
	}, nil
}
