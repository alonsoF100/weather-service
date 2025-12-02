package service

import (
	"log/slog"

	"github.com/alonsoF100/weather-service/internal/models"
)

type GeoClient interface {
	GetCordinates(city string) (*models.GeoLocation, error)
}

type WeatherClient interface {
	GetWeather(geoLocation *models.GeoLocation) (*models.WeatherData, error)
}

type Repository interface {
	GetWeather(city string) (*models.WeatherData, error)
	InsertWeather(weather *models.WeatherData) error
}

type WeatherService struct {
	geoClient     GeoClient
	weatherClient WeatherClient
	repository    Repository
}

func NewWeatherService(geoClient GeoClient, weatherClient WeatherClient, repository Repository) *WeatherService {
	return &WeatherService{
		geoClient:     geoClient,
		weatherClient: weatherClient,
		repository:    repository,
	}
}

func (s *WeatherService) GetWeatherByCity(city string) (*models.WeatherData, error) {
	slog.Info("Getting weather by city", "city", city)

	weatherData, err := s.repository.GetWeather(city)
	if err != nil {
		slog.Error("Failed to get weather from repository",
			"city", city,
			"error", err,
		)
		return nil, err
	}

	slog.Info("Successfully retrieved weather data",
		"city", city,
		"temperature", weatherData.Temperature,
	)

	return weatherData, nil
}

func (s *WeatherService) AddWeather(city string) error {
	slog.Info("Adding weather data for city", "city", city)

	slog.Debug("Getting coordinates for city", "city", city)
	cords, err := s.geoClient.GetCordinates(city)
	if err != nil {
		slog.Error("Failed to get coordinates",
			"city", city,
			"error", err,
		)
		return err
	}
	slog.Debug("Coordinates retrieved",
		"city", city,
		"lat", cords.Latitude,
		"lon", cords.Longitude,
	)

	slog.Debug("Getting weather data for coordinates",
		"lat", cords.Latitude,
		"lon", cords.Longitude,
	)
	weather, err := s.weatherClient.GetWeather(cords)
	if err != nil {
		slog.Error("Failed to get weather data",
			"city", city,
			"lat", cords.Latitude,
			"lon", cords.Longitude,
			"error", err,
		)
		return err
	}
	slog.Debug("Weather data retrieved",
		"city", city,
		"temperature", weather.Temperature,
	)

	slog.Debug("Saving weather data to database", "city", city)
	err = s.repository.InsertWeather(weather)
	if err != nil {
		slog.Error("Failed to save weather data to database",
			"city", city,
			"error", err,
		)
		return err
	}

	slog.Info("Weather data successfully added",
		"city", city,
		"temperature", weather.Temperature,
		"timestamp", weather.TimeStamp,
	)

	return nil
}
