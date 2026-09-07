package menu

import (
	"context"
	"log"

	"avito/internal/model"
)

func (r *Menu) Create(ctx context.Context, menuItem *model.MenuItem) error {
	query := `INSERT INTO menu_items
		(restaurant_id, name, description, price, category, is_available)
		VALUES (:restaurant_id, :name, :description, :price, :category, :is_available) RETURNING id, created_at, updated_at
		`

	rows, err := r.db.NamedQueryContext(ctx, query, menuItem)
	if err != nil {
		return err
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("close rows: %v", err)
		}
	}()

	if rows.Next() {
		return rows.StructScan(menuItem)
	}
	return nil
}
