package menu

import (
	"context"

	"avito/internal/model"
)

type Service interface {
	GetItemByID(ctx context.Context, id int) (*model.MenuItem, error)
	GetMenu(ctx context.Context, restaurantID int, filter *model.MenuFilter) ([]model.MenuItem, error)
	Create(ctx context.Context, menuItem *model.MenuItem) error
	Update(ctx context.Context, id int, updates map[string]any) (*model.MenuItem, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
