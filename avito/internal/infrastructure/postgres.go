package infrastructure

import (
	"avito/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

func NewPostgres(cfg *config.Postgres) (*sqlx.DB, error) {
	const maxConns = 25
	const maxIdleConns = 5
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdleConns)

	return db, nil
}
