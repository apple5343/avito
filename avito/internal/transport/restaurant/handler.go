package restaurant

import (
	"avito/internal/model"
	"context"
)

type RestaurantService interface {
	Create(ctx context.Context, restaurant *model.Restaurant) error
	Update(ctx context.Context, updates map[string]interface{}) error
}

type Handler struct {
	service RestaurantService
}

func NewHandler(service RestaurantService) *Handler {
	return &Handler{
		service: service,
	}
}
