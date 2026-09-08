package restaurant

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Restaurant) GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error) {
	query := `SELECT id, name, description, city, cuisine, callback_url, api_key, is_active, created_at, updated_at FROM restaurants WHERE api_key = $1`
	var restaurant model.Restaurant
	if err := r.db.GetContext(ctx, &restaurant, query, apiKey); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}

func (r *Restaurant) GetByID(ctx context.Context, id int) (*model.Restaurant, error) {
	query := `SELECT id, name, description, city, cuisine, callback_url, api_key, is_active, created_at, updated_at FROM restaurants WHERE id = $1`
	var restaurant model.Restaurant
	if err := r.db.GetContext(ctx, &restaurant, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}

func (r *Restaurant) List(ctx context.Context, filter *model.RestaurantFilter) ([]model.Restaurant, error) {
	query := `SELECT id, name, description, city, cuisine, is_active, created_at, updated_at 
			FROM restaurants`

	args := []any{}
	conditions := []string{}

	if filter.Name != nil && *filter.Name != "" {
		args = append(args, "%"+*filter.Name+"%")
		conditions = append(conditions, fmt.Sprintf("name ILIKE $%d", len(args)))
	}
	if filter.City != nil && *filter.City != "" {
		args = append(args, *filter.City)
		conditions = append(conditions, fmt.Sprintf("city = $%d", len(args)))
	}
	if filter.Cuisine != nil && *filter.Cuisine != "" {
		args = append(args, *filter.Cuisine)
		conditions = append(conditions, fmt.Sprintf("cuisine = $%d", len(args)))
	}
	if filter.IsActive != nil {
		args = append(args, *filter.IsActive)
		conditions = append(conditions, fmt.Sprintf("is_active = $%d", len(args)))
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY name ASC"

	var restaurants []model.Restaurant
	if err := r.db.SelectContext(ctx, &restaurants, query, args...); err != nil {
		return nil, err
	}
	return restaurants, nil
}
