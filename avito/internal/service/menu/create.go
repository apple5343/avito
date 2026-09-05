package menu

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
)

func (s *Menu) Create(ctx context.Context, menuItem *model.MenuItem) error {
	apiKey := ctxutil.APIKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}

	restaurant, err := s.restaurantService.GetByAPIKey(ctx, apiKey)
	if err != nil {
		return err
	}

	menuItem.RestaurantID = restaurant.ID

	if err = menuItem.BeforeCreate(); err != nil {
		return err
	}

	if err = s.repo.Create(ctx, menuItem); err != nil {
		return errorx.New(errorx.CodeInternalServerError, fmt.Errorf("create menu item: %w", err))
	}

	return nil
}
