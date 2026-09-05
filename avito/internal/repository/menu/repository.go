package menu

import "github.com/jmoiron/sqlx"

type Menu struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Menu {
	return &Menu{
		db: db,
	}
}
