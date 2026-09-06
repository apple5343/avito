package order

import (
	"context"

	"avito/internal/model"
)

type orderRepository interface {
	GetOrderByID(ctx context.Context, id int) (*model.Order, error)
	GetItemsByOrderID(ctx context.Context, orderID int) ([]model.OrderItem, error)
	GetUserOrders(ctx context.Context, userID int) ([]model.Order, error)
	Create(ctx context.Context, order *model.Order, items []model.OrderItem) error
	UpdateStatus(ctx context.Context, id int, status, message string) (*model.Order, error)
}

type menuService interface {
	GetItemByID(ctx context.Context, id int) (*model.MenuItem, error)
}

type restaurantService interface {
	GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error)
	GetByID(ctx context.Context, id int) (*model.Restaurant, error)
}

type restaurantAPI interface {
	Create(ctx context.Context, url string, order *model.Order, items []model.OrderItem) (bool, string, error)
}

type Order struct {
	repo              orderRepository
	menuService       menuService
	restaurantService restaurantService
	restaurant        restaurantAPI
}

func NewService(repo orderRepository, menuService menuService, restaurantAPI restaurantAPI, restaurantService restaurantService) *Order {
	return &Order{
		repo:              repo,
		menuService:       menuService,
		restaurantService: restaurantService,
		restaurant:        restaurantAPI,
	}
}
