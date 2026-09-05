package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID              int             `json:"id"`
	UserID          int             `json:"user_id"`
	RestaurantID    int             `json:"restaurant_id"`
	Status          string          `json:"status" db:"status"`
	DeliveryAddress string          `json:"delivery_address"`
	Comment         string          `json:"comment" db:"comment"`
	TotalPrice      decimal.Decimal `json:"total_price"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type OrderItem struct {
	ID         int             `json:"id"`
	OrderID    int             `json:"order_id"`
	MenuItemID int             `json:"menu_item_id"`
	Quantity   int             `json:"quantity"`
	Price      decimal.Decimal `json:"price"`
}
