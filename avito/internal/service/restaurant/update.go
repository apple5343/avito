package restaurant

import (
	"avito/internal/model"
	repoErrors "avito/internal/repository/errors"
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
	"avito/pkg/validator"
	"context"
	"errors"
	"fmt"
)

func (s *Restaurant) Update(ctx context.Context, updates map[string]interface{}) (*model.Restaurant, error) {
	apiKey := ctxutil.ApiKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return nil, errorx.New(errorx.CodeUnauthorized, fmt.Errorf("unauthorized"))
	}
	if len(updates) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("updates is required"))
	}

	if name, ok := updates["name"]; ok && len(name.(string)) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("name is required"))
	}
	if description, ok := updates["description"]; ok && len(description.(string)) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("description is required"))
	}
	if city, ok := updates["city"]; ok && len(city.(string)) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("city is required"))
	}
	if cuisine, ok := updates["cuisine"]; ok && len(cuisine.(string)) == 0 {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("cuisine is required"))
	}
	if callbackURL, ok := updates["callback_url"]; ok && !validator.ValidateURL(callbackURL.(string)) {
		return nil, errorx.New(errorx.CodeBadRequest, fmt.Errorf("callback url is invalid"))
	}

	restaurant, err := s.repo.Update(ctx, apiKey, updates)
	if err != nil {
		if errors.Is(err, repoErrors.ErrNotFound) {
			return nil, errorx.New(errorx.CodeNotFound, fmt.Errorf("restaurant not found"))
		}
		return nil, errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update restaurant: %w", err))
	}
	return restaurant, nil
}
