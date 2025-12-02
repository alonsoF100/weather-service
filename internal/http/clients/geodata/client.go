package geodata

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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

func (c *Client) GetCordinates(city string) (*models.GeoLocation, error) {
	baseURL := "https://geocoding-api.open-meteo.com/v1/search"

	params := url.Values{}
	params.Add("name", city)
	params.Add("count", "1")

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
		Results []struct {
			Name      string  `json:"name"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Timezone  string  `json:"timezone"`
		} `json:"results"`
	}

	if err := json.NewDecoder(response.Body).Decode(&temp); err != nil {
		return nil, err
	}

	return &models.GeoLocation{
		Name:      temp.Results[0].Name,
		Latitude:  temp.Results[0].Latitude,
		Longitude: temp.Results[0].Longitude,
		Timezone:  temp.Results[0].Timezone,
	}, nil
}
