package restaurant

import (
	"avito/internal/model"
	"context"
)

func (r *Restaurant) Create(ctx context.Context, restaurant *model.Restaurant) error {
	query := `INSERT INTO restaurants
		(name, description, city, cuisine, callback_url, api_key)
		VALUES (:name, :description, :city, :cuisine, :callback_url, :api_key) RETURNING id, created_at, updated_at
		`

	rows, err := r.db.NamedQueryContext(ctx, query, restaurant)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return rows.StructScan(restaurant)
	}
	return nil
}
