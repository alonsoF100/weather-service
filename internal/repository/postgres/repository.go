package postgres

import (
	"context"
	"log/slog"

	"github.com/alonsoF100/weather-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetWeather(city string) (*models.WeatherData, error) {
	const query = `
	SELECT name, temperature, time_stamp FROM readings
	WHERE name ILIKE $1
	ORDER BY time_stamp DESC
	LIMIT 1`

	slog.Debug("Getting weather from database",
		"city", city,
		"query", query,
	)

	var weather models.WeatherData
	err := r.pool.QueryRow(context.Background(), query, city).Scan(
		&weather.Name,
		&weather.Temperature,
		&weather.TimeStamp)
	if err != nil {
		slog.Error("Failed to get weather from database",
			"city", city,
			"error", err,
		)
		return nil, err
	}

	slog.Debug("Successfully retrieved weather data",
		"city", city,
		"temperature", weather.Temperature,
		"timestamp", weather.TimeStamp,
	)

	return &weather, nil
}

func (r *Repository) InsertWeather(weather *models.WeatherData) error {
	const query = `
	INSERT INTO readings (name, temperature, time_stamp)
	VALUES ($1, $2, $3)`

	slog.Debug("Inserting weather into database",
		"city", weather.Name,
		"temperature", weather.Temperature,
		"timestamp", weather.TimeStamp,
		"query", query,
	)

	_, err := r.pool.Exec(context.Background(), query, weather.Name, weather.Temperature, weather.TimeStamp)
	if err != nil {
		slog.Error("Failed to insert weather into database",
			"city", weather.Name,
			"error", err,
		)
		return err
	}

	slog.Info("Weather data inserted successfully",
		"city", weather.Name,
		"temperature", weather.Temperature,
	)

	return nil
}
