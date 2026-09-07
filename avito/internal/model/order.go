package model

import (
	"errors"
	"time"

	"avito/pkg/errorx"

	"github.com/shopspring/decimal"
)

const (
	StatusPending    = "pending"
	StatusConfirmed  = "confirmed"
	StatusRejected   = "rejected"
	StatusCooking    = "cooking"
	StatusReady      = "ready"
	StatusDelivering = "delivering"
	StatusDelivered  = "delivered"
)

var Statuses = map[string]bool{
	"pending":    false,
	"confirmed":  true,
	"rejected":   true,
	"cooking":    true,
	"ready":      true,
	"delivering": true,
	"delivered":  true,
}

type Order struct {
	ID                int             `json:"id" db:"id"`
	UserID            int             `json:"user_id" db:"user_id"`
	RestaurantID      int             `json:"restaurant_id" db:"restaurant_id"`
	Status            string          `json:"status" db:"status"`
	DeliveryAddress   string          `json:"delivery_address" db:"delivery_address"`
	Comment           string          `json:"comment" db:"comment"`
	TotalPrice        decimal.Decimal `json:"total_price" db:"total_price"`
	RestaurantMessage string          `json:"restaurant_message,omitempty" db:"restaurant_message"`
	CreatedAt         time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" db:"updated_at"`
}

func (o *Order) BeforeCreate() error {
	if len(o.DeliveryAddress) == 0 {
		return errorx.New(errorx.CodeBadRequest, errors.New("delivery address is required"))
	}
	o.Status = StatusPending
	return nil
}

type OrderItem struct {
	ID         int             `json:"id" db:"id"`
	OrderID    int             `json:"order_id" db:"order_id"`
	MenuItemID int             `json:"menu_item_id" db:"menu_item_id"`
	Quantity   int             `json:"quantity" db:"quantity"`
	Price      decimal.Decimal `json:"price" db:"price"`
}

type OrderRequest struct {
	RestaurantID    int                `json:"restaurant_id" db:"restaurant_id"`
	DeliveryAddress string             `json:"delivery_address" db:"delivery_address"`
	Comment         string             `json:"comment" db:"comment"`
	Items           []OrderItemRequest `json:"items" db:"-"`
}

type OrderItemRequest struct {
	MenuItemID int `json:"menu_item_id" db:"menu_item_id"`
	Quantity   int `json:"quantity" db:"quantity"`
}

type OrderHistory struct {
	ID        int       `json:"id" db:"id"`
	OrderID   int       `json:"order_id" db:"order_id"`
	Status    string    `json:"status" db:"status"`
	ChangedAt time.Time `json:"changed_at" db:"changed_at"`
}
