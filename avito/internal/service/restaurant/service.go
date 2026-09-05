package restaurant

import (
	"context"

	"avito/internal/model"
)

type restaurantRepository interface {
	GetByID(ctx context.Context, id int) (*model.Restaurant, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error)
	List(ctx context.Context, filter *model.RestaurantFilter) ([]model.Restaurant, error)
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
