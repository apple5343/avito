package restaurant

import "github.com/jmoiron/sqlx"

type Restaurant struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Restaurant {
	return &Restaurant{
		db: db,
	}
}
