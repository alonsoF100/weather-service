package main

import (
	"log"
	"log/slog"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/alonsoF100/weather-service/internal/logger"
)

func main() {
	config := config.Load()
	log.Printf("config loaded successfully")

	logger.Setup(config.Logger)
	slog.Info("logger setup successfully")
}
