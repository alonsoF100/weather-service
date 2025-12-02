package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	weatherData, err := h.weatherService.GetWeatherByCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		slog.Error("Failed encode to json", "error", err)
		return
	}
}
