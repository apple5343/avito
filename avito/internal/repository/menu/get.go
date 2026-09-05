package menu

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

func (r *Menu) GetMenu(ctx context.Context, restaurantID int, filter *model.MenuFilter) ([]model.MenuItem, error) {
	query := `SELECT id, name, description, price, category, is_available, created_at, updated_at 
			FROM menu_items`

	args := []any{restaurantID}
	conditions := []string{"restaurant_id = $1"}

	if filter.Category != nil && *filter.Category != "" {
		args = append(args, *filter.Category)
		conditions = append(conditions, fmt.Sprintf("category = $%d", len(args)))
	}
	if filter.IsAvailable != nil {
		args = append(args, *filter.IsAvailable)
		conditions = append(conditions, fmt.Sprintf("is_available = $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY name ASC"

	var menuItems []model.MenuItem
	if err := r.db.SelectContext(ctx, &menuItems, query, args...); err != nil {
		return nil, err
	}
	return menuItems, nil
}
