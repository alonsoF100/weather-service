package handler

import (
	"github.com/alonsoF100/weather-service/internal/models"
)

type WeatherService interface {
	GetWeatherByCity(city string) (*models.WeatherData, error)
}

type Handler struct {
	weatherService WeatherService
}

func New(weatherService WeatherService) *Handler {
	return &Handler{
		weatherService: weatherService,
	}
}
