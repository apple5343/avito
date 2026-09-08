package restaurant

import (
	"context"
	"log"

	"avito/internal/model"
)

func (r *Restaurant) Create(ctx context.Context, restaurant *model.Restaurant) error {
	query := `INSERT INTO restaurants
		(name, description, city, cuisine, is_active, callback_url, api_key)
		VALUES (:name, :description, :city, :cuisine, :is_active, :callback_url, :api_key) RETURNING id, created_at, updated_at
		`

	rows, err := r.db.NamedQueryContext(ctx, query, restaurant)
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
		return rows.StructScan(restaurant)
	}
	return nil
}
