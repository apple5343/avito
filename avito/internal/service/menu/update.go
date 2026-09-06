package menu

import (
	"context"
	"errors"

	"avito/internal/model"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
)

func (s *Menu) Update(ctx context.Context, id int, updates map[string]any) (*model.MenuItem, error) {
	apiKey := ctxutil.APIKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}
	if len(updates) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, errors.New("updates is required"))
	}

	restaurant, err := s.restaurantService.GetByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	menuItem, err := s.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if restaurant.ID != menuItem.RestaurantID {
		return nil, errorx.New(errorx.CodeForbidden, errors.New("forbidden"))
	}

	if name, ok := updates["name"]; ok {
		if nameStr, okk := name.(string); okk && len(nameStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("name is required"))
		}
	}
	if description, ok := updates["description"]; ok {
		if descriptionStr, okk := description.(string); okk && len(descriptionStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("description is required"))
		}
	}
	if price, ok := updates["price"]; ok {
		if priceStr, okk := price.(string); okk && len(priceStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("price is required"))
		}
	}
	if category, ok := updates["category"]; ok {
		if categoryStr, okk := category.(string); okk && len(categoryStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("category is required"))
		}
	}
	menuItem, err = s.repo.Update(ctx, menuItem.ID, updates)
	if err != nil {
		return nil, err
	}
	return menuItem, nil
}
