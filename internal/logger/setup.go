package logger

import (
	"log/slog"
	"os"

	"github.com/alonsoF100/weather-service/internal/config"
)

func Setup(cfg config.LoggerConfig) *slog.Logger {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: parse(cfg.Level)})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("Logger setuped successfully")
	return logger
}

func parse(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
