package order

import (
	"context"
	"database/sql"
	"errors"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Order) GetOrderByID(ctx context.Context, id int) (*model.Order, error) {
	order := model.Order{}
	query := `SELECT id, user_id, restaurant_id, status, delivery_address, comment, total_price, restaurant_message, created_at, updated_at 
			FROM orders WHERE id = $1`
	if err := r.db.GetContext(ctx, &order, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *Order) GetItemsByOrderID(ctx context.Context, orderID int) ([]model.OrderItem, error) {
	var items []model.OrderItem
	query := `SELECT id, order_id, menu_item_id, quantity, price
			FROM order_items WHERE order_id = $1
			ORDER BY id`
	if err := r.db.SelectContext(ctx, &items, query, orderID); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Order) GetUserOrders(ctx context.Context, userID int) ([]model.Order, error) {
	var orders []model.Order
	query := `SELECT id, user_id, restaurant_id, status, delivery_address, comment, total_price, restaurant_message, created_at, updated_at 
			FROM orders WHERE user_id = $1
			ORDER BY created_at DESC`
	if err := r.db.SelectContext(ctx, &orders, query, userID); err != nil {
		return nil, err
	}
	return orders, nil
}
