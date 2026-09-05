package restaurant

import (
	"avito/internal/model"
	"context"
)

type restaurantRepository interface {
	Create(ctx context.Context, restaurant *model.Restaurant) error
	Update(ctx context.Context, apiKey string, updates map[string]interface{}) (*model.Restaurant, error)
}

type Restaurant struct {
	repo restaurantRepository
}

func NewService(repo restaurantRepository) *Restaurant {
	return &Restaurant{
		repo: repo,
	}
}
