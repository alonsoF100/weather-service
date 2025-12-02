package cron

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/go-co-op/gocron/v2"
)

type WeatherService interface {
	AddWeather(city string) error
}

type Scheduler struct {
	weatherService WeatherService
	scheduler      gocron.Scheduler
	cfg            config.CronConfig
}

func NewScheduler(weatherService WeatherService, cfg config.CronConfig) *Scheduler {
	return &Scheduler{
		weatherService: weatherService,
		cfg:            cfg,
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	slog.Info("Starting cron scheduler")

	sched, err := gocron.NewScheduler()
	if err != nil {
		return fmt.Errorf("failed to create scheduler: %w", err)
	}
	s.scheduler = sched

	_, err = sched.NewJob(
		gocron.DurationJob(s.cfg.UpdateInterval),
		gocron.NewTask(
			func() {
				startTime := time.Now()
				slog.Info("Cron job started",
					"start_time", startTime.Format("15:04:05"),
				)

				cities := s.cfg.Cities

				for _, city := range cities {
					cityStart := time.Now()
					slog.Info("Processing city",
						"city", city,
						"time", cityStart.Format("15:04:05"),
					)

					if err := s.weatherService.AddWeather(city); err != nil {
						slog.Error("Failed to add weather",
							"city", city,
							"error", err,
							"duration_ms", time.Since(cityStart).Milliseconds(),
						)
					} else {
						slog.Info("Weather data updated",
							"city", city,
							"duration_ms", time.Since(cityStart).Milliseconds(),
						)
					}

					time.Sleep(100 * time.Millisecond)
				}

				slog.Info("Cron job completed",
					"total_duration_ms", time.Since(startTime).Milliseconds(),
					"end_time", time.Now().Format("15:04:05\n"),
				)
			},
		),
		gocron.WithSingletonMode(gocron.LimitModeWait),
	)
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	sched.Start()
	slog.Info("Cron scheduler started successfully")

	<-ctx.Done()
	slog.Info("Stopping cron scheduler")
	return s.scheduler.Shutdown()
}
