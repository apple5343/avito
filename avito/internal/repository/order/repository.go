package order

import "github.com/jmoiron/sqlx"

type Order struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Order {
	return &Order{
		db: db,
	}
}
