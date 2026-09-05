package menu

import (
	"context"
	"database/sql"
	"errors"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Menu) GetItemByID(ctx context.Context, id int) (*model.MenuItem, error) {
	query := "SELECT id, restaurant_id, name, description, price, category, is_available, created_at, updated_at FROM menu_items WHERE id = $1"
	menuItem := &model.MenuItem{}
	if err := r.db.GetContext(ctx, menuItem, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, err
	}
	return menuItem, nil
}
