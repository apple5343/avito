package restaurant

import (
	"avito/internal/util/ctxutil"
	"avito/pkg/errorx"
	"context"
	"fmt"
)

func (s *Restaurant) Update(ctx context.Context, updates map[string]interface{}) error {
	apiKey := ctxutil.ApiKeyFromContext(ctx)
	if len(apiKey) == 0 {
		return errorx.New(errorx.CodeUnauthorized, fmt.Errorf("unauthorized"))
	}

	if city, ok := updates["city"]; ok && len(city.(string)) == 0 {
		return errorx.New(errorx.CodeBadRequest, fmt.Errorf("city is required"))
	}
	if cuisine, ok := updates["cuisine"]; ok && len(cuisine.(string)) == 0 {
		return errorx.New(errorx.CodeBadRequest, fmt.Errorf("cuisine is required"))
	}
	if callbackURL, ok := updates["callback_url"]; ok && len(callbackURL.(string)) == 0 {
		return errorx.New(errorx.CodeBadRequest, fmt.Errorf("callback url is required"))
	}
	
	if err := s.repo.Update(ctx, updates); err != nil {
		return errorx.New(errorx.CodeInternalServerError, fmt.Errorf("update restaurant: %w", err))

	}
	return nil
}
