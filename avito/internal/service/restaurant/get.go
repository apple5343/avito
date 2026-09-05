package restaurant

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"avito/pkg/errorx"
)

func (s *Restaurant) GetByAPIKey(ctx context.Context, apiKey string) (*model.Restaurant, error) {
	restaurant, err := s.repo.GetByAPIKey(ctx, apiKey)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, fmt.Errorf("restaurant not found: %w", err))
		}
		return nil, err
	}
	return restaurant, nil
}

func (s *Restaurant) GetByID(ctx context.Context, id int) (*model.Restaurant, error) {
	restaurant, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, fmt.Errorf("restaurant not found: %w", err))
		}
		return nil, err
	}
	return restaurant, nil
}

func (s *Restaurant) List(ctx context.Context, filter *model.RestaurantFilter) ([]model.Restaurant, error) {
	if filter.IsActive == nil {
		active := true
		filter.IsActive = &active
	}
	restaurants, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("list restaurants: %w", err))
	}

	return restaurants, nil
}
