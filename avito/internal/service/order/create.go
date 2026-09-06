package order

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"

	"github.com/shopspring/decimal"
)

func (s *Order) Create(ctx context.Context, req *model.OrderRequest) (*model.Order, []model.OrderItem, error) {
	userID := ctxutil.UserIDFromContext(ctx)
	if userID == 0 {
		return nil, nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}

	if len(req.Items) == 0 {
		return nil, nil, errorx.New(errorx.CodeBadRequest, errors.New("items is required"))
	}

	for _, reqItem := range req.Items {
		if reqItem.Quantity <= 0 {
			return nil, nil, errorx.New(errorx.CodeBadRequest, errors.New("quantity must be greater than 0"))
		}
	}

	restaurant, err := s.restaurantService.GetByID(ctx, req.RestaurantID)
	if err != nil {
		return nil, nil, err
	}

	if !restaurant.IsActive {
		return nil, nil, errorx.New(errorx.CodeBadRequest, errors.New("restaurant is not available"))
	}

	order := &model.Order{
		UserID:          userID,
		RestaurantID:    req.RestaurantID,
		DeliveryAddress: req.DeliveryAddress,
		Comment:         req.Comment,
	}

	if err = order.BeforeCreate(); err != nil {
		return nil, nil, err
	}

	totalPrice := decimal.NewFromInt(0)
	items := make([]model.OrderItem, 0, len(req.Items))
	for _, reqItem := range req.Items {
		item, err := s.menuService.GetItemByID(ctx, reqItem.MenuItemID)
		if err != nil {
			return nil, nil, err
		}
		if item.RestaurantID != order.RestaurantID {
			return nil, nil, errorx.New(errorx.CodeBadRequest, errors.New("invalid menu item"))
		}
		if !item.IsAvailable {
			return nil, nil, errorx.New(errorx.CodeBadRequest, errors.New("menu item is not available"))
		}

		items = append(items, model.OrderItem{
			MenuItemID: reqItem.MenuItemID,
			Quantity:   reqItem.Quantity,
			Price:      item.Price,
		})

		totalPrice = totalPrice.Add(item.Price.Mul(decimal.NewFromInt(int64(reqItem.Quantity))))
	}

	order.TotalPrice = totalPrice

	if err = s.repo.Create(ctx, order, items); err != nil {
		return nil, nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("create order: %w", err))
	}

	status, message, err := s.restaurant.Create(ctx, restaurant.CallbackURL, order, items)
	if err != nil {
		return nil, nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("notify restaurant: %w", err))
	}
	if message == "" && !status {
		message = "Ресторан не принял заказ"
	} else if message == "" && status {
		message = "Ресторан принял заказ"
	}

	if !status {
		if _, err = s.repo.UpdateStatus(ctx, order.ID, model.StatusCancelled, message); err != nil {
			return nil, nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update order status: %w", err))
		}
		return nil, nil, errorx.New(errorx.CodeConflict, errors.New(message))
	}

	if _, err = s.repo.UpdateStatus(ctx, order.ID, model.StatusConfirmed, message); err != nil {
		return nil, nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update order status: %w", err))
	}
	order.Status = model.StatusConfirmed

	return order, items, nil
}
