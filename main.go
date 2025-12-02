package main

import (
	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/alonsoF100/weather-service/internal/http/server"
	"github.com/alonsoF100/weather-service/internal/logger"
)

func main() {
	config := config.Load()

	logger.Setup(config.Logger)

	server.Setup(config.Server).Start()
}
