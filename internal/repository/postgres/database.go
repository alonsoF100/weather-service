package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/alonsoF100/weather-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func NewPool(cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSlMode,
	)
	slog.Info("Connection string:", "connstring", connString)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	connConfig := poolConfig.ConnConfig
	db := stdlib.OpenDB(*connConfig)

	if err := goose.Up(db, "./migrations/postgres"); err != nil {
		return nil, err
	}

	return pool, nil
}
