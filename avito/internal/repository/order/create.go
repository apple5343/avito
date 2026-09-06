package order

import (
	"context"
	"fmt"
	"strings"

	"avito/internal/model"

	"log"
)

func (r *Order) Create(ctx context.Context, order *model.Order, items []model.OrderItem) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("rollback transaction: %v", rbErr)
			}
		}
	}()

	orderQuery := `INSERT INTO orders (user_id, restaurant_id, status, delivery_address, comment, total_price, restaurant_message)
			VALUES (:user_id, :restaurant_id, :status, :delivery_address, :comment, :total_price, :restaurant_message) 
			RETURNING id, created_at, updated_at`

	rows, err := tx.NamedQuery(orderQuery, order)
	if err != nil {
		return fmt.Errorf("insert order: %w", err)
	}

	if rows.Next() {
		if err = rows.StructScan(order); err != nil {
			rows.Close()
			return fmt.Errorf("scan order: %w", err)
		}
	}
	rows.Close()

	if order.ID == 0 {
		return fmt.Errorf("failed to get order id")
	}

	if len(items) == 0 {
		return nil
	}

	for i := range items {
		items[i].OrderID = order.ID
	}

	values := make([]string, 0, len(items))
	args := make([]any, 0, len(items)*4)
	argCounter := 1
	for _, item := range items {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d)", argCounter, argCounter+1, argCounter+2, argCounter+3))
		args = append(args, item.OrderID, item.MenuItemID, item.Quantity, item.Price)
		argCounter += 4
	}
	itemQuery := fmt.Sprintf(
		`INSERT INTO order_items (order_id, menu_item_id, quantity, price)
			VALUES %s RETURNING id`,
		strings.Join(values, ", "),
	)
	rows, err = tx.QueryxContext(ctx, itemQuery, args...)
	if err != nil {
		return fmt.Errorf("insert order items: %w", err)
	}
	defer rows.Close()
	idx := 0
	for rows.Next() {
		if idx < len(items) {
			if err = rows.StructScan(&items[idx]); err != nil {
				return fmt.Errorf("scan order item: %w", err)
			}
		} else {
			return fmt.Errorf("too many order items")
		}
		idx++
	}

	if idx != len(items) {
		return fmt.Errorf("expected %d order items, got %d", len(items), idx)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
