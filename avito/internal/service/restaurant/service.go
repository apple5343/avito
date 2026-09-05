package restaurant

import (
	"context"

	"avito/internal/model"
)

type restaurantRepository interface {
	GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error)
	Create(ctx context.Context, restaurant *model.Restaurant) error
	Update(ctx context.Context, apiKey string, updates map[string]any) (*model.Restaurant, error)
}

type Restaurant struct {
	repo restaurantRepository
}

func NewService(repo restaurantRepository) *Restaurant {
	return &Restaurant{
		repo: repo,
	}
}
