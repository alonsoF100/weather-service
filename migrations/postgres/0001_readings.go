package postgres

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateReadings, downCreateReadings)
}

func upCreateReadings(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE readings (
			id BIGSERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			temperature DECIMAL(5,2) NOT NULL,
			time_stamp TIMESTAMPTZ NOT NULL
		);
	`)
	return err
}

func downCreateReadings(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE readings;")
	return err
}
