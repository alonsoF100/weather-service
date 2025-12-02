package main

import (
	"context"
	"log/slog"
	"sync"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/alonsoF100/weather-service/internal/cron"
	"github.com/alonsoF100/weather-service/internal/http/clients/geodata"
	"github.com/alonsoF100/weather-service/internal/http/clients/weather"
	"github.com/alonsoF100/weather-service/internal/http/handler"
	"github.com/alonsoF100/weather-service/internal/http/server"
	"github.com/alonsoF100/weather-service/internal/logger"
	"github.com/alonsoF100/weather-service/internal/repository/postgres"
	"github.com/alonsoF100/weather-service/internal/service"
	_ "github.com/alonsoF100/weather-service/migrations/postgres"
)

func main() {
	config := config.Load()

	logger.Setup(config.Logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	geoClient := geodata.New(config.Client)
	weatherClient := weather.New(config.Client)

	pool, err := postgres.NewPool(config.Database)
	if err != nil {
		slog.Error("Failed to create pool", "error", err)
		panic(err)
	}
	defer pool.Close()

	dataBase := postgres.New(pool)

	weatherService := service.NewWeatherService(geoClient, weatherClient, dataBase)

	handler := handler.New(weatherService)
	server := server.Setup(config.Server, handler)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		server.Start()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()

		scheduler := cron.NewScheduler(weatherService, config.Cron)
		if err := scheduler.Start(ctx); err != nil {
			slog.Error("Cron scheduler failed", "error", err)
		}
	}()

	wg.Wait()
	slog.Info("Application stopped")
}
