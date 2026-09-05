package restaurant

import (
	"encoding/json"
	"fmt"
	"net/http"

	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Update() echo.HandlerFunc {
	allowedFields := map[string]bool{
		"name":         true,
		"description":  true,
		"city":         true,
		"cuisine":      true,
		"callback_url": true,
		"is_active":    true,
	}
	return func(c echo.Context) error {
		var updates map[string]any
		if err := json.NewDecoder(c.Request().Body).Decode(&updates); err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid json: %w", err))
		}

		for field := range updates {
			if !allowedFields[field] {
				return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid field: %s", field))
			}
		}

		restaurant, err := h.service.Update(c.Request().Context(), updates)

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, restaurant)
	}
}
