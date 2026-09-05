package model

import (
	"errors"
	"time"

	"avito/pkg/errorx"

	"github.com/shopspring/decimal"
)

type MenuItem struct {
	ID           int             `json:"id" db:"id"`
	RestaurantID int             `json:"restaurant_id" db:"restaurant_id"`
	Name         string          `json:"name" db:"name"`
	Description  string          `json:"description" db:"description"`
	Price        decimal.Decimal `json:"price" db:"price"`
	Category     string          `json:"category" db:"category"`
	IsAvailable  bool            `json:"is_available" db:"is_available"`
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
}

func (m *MenuItem) BeforeCreate() error {
	if len(m.Name) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("name is required"))
	}
	if m.Price.LessThan(decimal.Zero) {
		return errorx.New(errorx.CodeBadRequest, errors.New("price must be greater than 0"))
	}
	if len(m.Category) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("category is required"))
	}
	return nil
}
