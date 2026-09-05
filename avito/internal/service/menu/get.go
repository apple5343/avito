package menu

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"avito/pkg/errorx"
)

func (s *Menu) GetItemByID(ctx context.Context, id int) (*model.MenuItem, error) {
	item, err := s.repo.GetItemByID(ctx, id)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, fmt.Errorf("menu item not found: %w", err))
		}
		return nil, err
	}

	return item, nil
}

func (s *Menu) GetMenu(ctx context.Context, restaurantID int, filter *model.MenuFilter) ([]model.MenuItem, error) {
	if filter.IsAvailable == nil {
		available := true
		filter.IsAvailable = &available
	}
	items, err := s.repo.GetMenu(ctx, restaurantID, filter)
	if err != nil {
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("get menu: %w", err))
	}

	return items, nil
}
