package order

import (
	"context"
	"time"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Order) UpdateStatus(ctx context.Context, orderID int, status string) (*model.Order, error) {
	query := `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`
	res, err := r.db.ExecContext(ctx, query, status, time.Now(), orderID)
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
	return r.GetOrderByID(ctx, orderID)
}
