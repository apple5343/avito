package order

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
)

func (s *Order) GetOrderByID(ctx context.Context, id int) (*model.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, errors.New("order not found"))
		}
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("get order by id: %w", err))
	}
	return order, nil
}

func (s *Order) GetItemsByOrderID(ctx context.Context, orderID int) ([]model.OrderItem, error) {
	items, err := s.repo.GetItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("get items by order id: %w", err))
	}
	return items, nil
}

func (s *Order) GetUserOrder(ctx context.Context, orderID int) (*model.Order, []model.OrderItem, error) {
	userID := ctxutil.UserIDFromContext(ctx)
	if userID == 0 {
		return nil, nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}
	order, err := s.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, nil, err
	}
	if order.UserID != userID {
		return nil, nil, errorx.New(errorx.CodeForbidden, errors.New("forbidden"))
	}
	items, err := s.repo.GetItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("get items by order id: %w", err))
	}
	return order, items, nil
}

func (s *Order) GetUserOrders(ctx context.Context) ([]model.Order, error) {
	userID := ctxutil.UserIDFromContext(ctx)
	if userID == 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}
	orders, err := s.repo.GetUserOrders(ctx, userID)
	if err != nil {
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("get user orders: %w", err))
	}
	return orders, nil
}
