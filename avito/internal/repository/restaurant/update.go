package restaurant

import (
	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"context"
	"fmt"
	"strings"
	"time"
)

func (r *Restaurant) Update(ctx context.Context, apiKey string, updates map[string]interface{}) (*model.Restaurant, error) {
	setClauses := make([]string, 0, len(updates))
	args := []interface{}{}

	for field, value := range updates {
		args = append(args, value)
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", field, len(args)))
	}

	args = append(args, time.Now())
	setClauses = append(setClauses, fmt.Sprintf("updated_at = $%d", len(args)))

	args = append(args, apiKey)
	query := fmt.Sprintf("UPDATE restaurants SET %s WHERE api_key = $%d", strings.Join(setClauses, ", "), len(args))
	res, err := r.db.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, repoErrors.ErrNotFound
	}
	return r.GetByApiKey(ctx, apiKey)
}
