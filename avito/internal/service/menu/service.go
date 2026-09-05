package menu

import (
	"context"

	"avito/internal/model"
)

type menuRepository interface {
	GetItemByID(ctx context.Context, id int) (*model.MenuItem, error)
	Create(ctx context.Context, menuItem *model.MenuItem) error
	Update(ctx context.Context, id int, updates map[string]any) (*model.MenuItem, error)
}

type restaurantService interface {
	GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error)
}

type Menu struct {
	repo              menuRepository
	restaurantService restaurantService
}

func NewService(repo menuRepository, restaurantService restaurantService) *Menu {
	return &Menu{
		repo:              repo,
		restaurantService: restaurantService,
	}
}
