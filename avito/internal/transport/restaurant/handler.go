package restaurant

import (
	"context"

	"avito/internal/model"
)

type Service interface {
	Create(ctx context.Context, restaurant *model.Restaurant) error
	Update(ctx context.Context, updates map[string]any) (*model.Restaurant, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
