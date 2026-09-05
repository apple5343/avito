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
