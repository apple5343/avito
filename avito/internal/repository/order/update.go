package order

import (
	"context"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
)

func (r *Order) UpdateStatus(ctx context.Context, orderID int, status string) (*model.Order, error) {
	query := `UPDATE orders SET status = $2 WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, orderID, status)
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
