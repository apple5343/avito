package restaurant

import (
	"context"
	"fmt"

	"avito/internal/model"
	"avito/pkg/errorx"
)

func (s *Restaurant) Create(ctx context.Context, restaurant *model.Restaurant) error {
	if err := restaurant.BeforeCreate(); err != nil {
		return err
	}

	if err := s.repo.Create(ctx, restaurant); err != nil {
		return errorx.New(errorx.CodeInternalServerError, fmt.Errorf("create restaurant: %w", err))
	}

	return nil
}
