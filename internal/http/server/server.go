package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/alonsoF100/weather-service/internal/http/handler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func Setup(cfg config.ServerConfig, h *handler.Handler) *Server {
	r := chi.NewRouter()

	r.Get("/{city}", h.GetWeather)

	httpServer := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		server: &httpServer,
	}
}

func (server *Server) Start() {
	slog.Info("Starting server", "address", server.server.Addr)

	if err := server.server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
