package restaurant

import (
	"context"

	"avito/internal/model"
)

type Service interface {
	GetByID(ctx context.Context, id int) (*model.Restaurant, error)
	List(ctx context.Context, filter *model.RestaurantFilter) ([]model.Restaurant, error)
	Create(ctx context.Context, restaurant *model.Restaurant) error
	Update(ctx context.Context, restaurantID int, updates map[string]any) (*model.Restaurant, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
