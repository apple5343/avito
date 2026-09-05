package mock

import (
	"context"

	"avito/internal/model"
)

type RestaurantMock struct{}

func New() *RestaurantMock {
	return &RestaurantMock{}
}

func (r *RestaurantMock) Create(ctx context.Context, url string, order *model.Order, items []model.OrderItem) (bool, string, error) {
	return true, "", nil
}
