package order

import (
	"context"

	"avito/internal/model"
)

type Service interface {
	GetUserOrder(ctx context.Context, orderID int) (*model.Order, []model.OrderItem, error)
	GetUserOrders(ctx context.Context) ([]model.Order, error)
	Create(ctx context.Context, req *model.OrderRequest) (*model.Order, []model.OrderItem, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
