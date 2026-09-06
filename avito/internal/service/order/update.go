package order

import (
	"context"
	"errors"
	"fmt"
	"log"

	"avito/internal/model"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
)

func (s *Order) UpdateStatus(ctx context.Context, orderID int, status, message string) (*model.Order, error) {
	apiKey := ctxutil.APIKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}
	if !model.Statuses[status] {
		return nil, errorx.New(errorx.CodeBadRequest, errors.New("invalid status"))
	}
	restaurant, err := s.restaurantService.GetByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}
	order, err := s.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order.RestaurantID != restaurant.ID {
		return nil, errorx.New(errorx.CodeForbidden, errors.New("forbidden"))
	}

	if order.Status == model.StatusCancelled || order.Status == model.StatusDelivered {
		return nil, errorx.New(errorx.CodeBadRequest, errors.New("order is already cancelled or delivered"))
	}
	order, err = s.repo.UpdateStatus(ctx, orderID, status, message)
	if err != nil {
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update order status: %w", err))
	}
	log.Printf("order %d status updated to %s", orderID, status)
	return order, nil
}
