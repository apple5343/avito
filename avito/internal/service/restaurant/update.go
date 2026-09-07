package restaurant

import (
	"context"
	"errors"
	"fmt"

	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
	"avito/pkg/validator"
)

func (s *Restaurant) Update(ctx context.Context, restaurantID int, updates map[string]any) (*model.Restaurant, error) {
	apiKey := ctxutil.APIKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, errors.New("unauthorized"))
	}
	restaurant, err := s.GetByID(ctx, restaurantID)
	if err != nil {
		return nil, err
	}
	if restaurant.APIKey != apiKey {
		return nil, errorx.New(errorx.CodeForbidden, errors.New("forbidden"))
	}
	if len(updates) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, errors.New("updates is required"))
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
	if city, ok := updates["city"]; ok {
		if cityStr, okk := city.(string); okk && len(cityStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("city is required"))
		}
	}
	if cuisine, ok := updates["cuisine"]; ok {
		cusineStr, okk := cuisine.(string)
		if okk && len(cusineStr) == 0 {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("cuisine is required"))
		}
		if !model.Cuisines[cusineStr] {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("invalid cuisine type"))
		}
	}
	if callbackURL, ok := updates["callback_url"]; ok {
		if callbackURLStr, okk := callbackURL.(string); okk && !validator.ValidateURL(callbackURLStr) {
			return nil, errorx.New(errorx.CodeBadRequest, errors.New("callback url is required"))
		}
	}

	restaurant, err = s.repo.Update(ctx, restaurantID, updates)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, errors.New("restaurant not found"))
		}
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update restaurant: %w", err))
	}
	return restaurant, nil
}
