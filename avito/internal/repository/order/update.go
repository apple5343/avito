package order

import (
	"context"
	"database/sql"
	"time"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Order) UpdateStatus(ctx context.Context, orderID int, status string, message string) (*model.Order, error) {
	var res sql.Result
	var err error
	if message != "" {
		query := `UPDATE orders SET status = $1, restaurant_message = $2, updated_at = $3 WHERE id = $4`
		res, err = r.db.ExecContext(ctx, query, status, message, time.Now(), orderID)
	} else {
		query := `UPDATE orders SET status = $1, updated_at = $2 WHERE id = $3`
		res, err = r.db.ExecContext(ctx, query, status, time.Now(), orderID)
	}
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
