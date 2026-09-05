package menu

import (
	"context"

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
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(menuItem)
	}
	return nil
}
