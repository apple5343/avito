package restaurant

import (
	"context"
	"database/sql"
	"errors"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Restaurant) GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error) {
	query := `SELECT id, name, description, city, cuisine, callback_url, api_key, created_at, updated_at FROM restaurants WHERE api_key = $1`
	var restaurant model.Restaurant
	if err := r.db.GetContext(ctx, &restaurant, query, apiKey); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repoErrors.ErrNotFound
		}
		return nil, err
	}
	return &restaurant, nil
}
