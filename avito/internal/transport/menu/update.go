package menu

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Update() echo.HandlerFunc {
	allowedFields := map[string]bool{
		"name":         true,
		"description":  true,
		"price":        true,
		"category":     true,
		"is_available": true,
	}
	return func(c echo.Context) error {
		idStr := c.Param("id")
		if len(idStr) == 0 {
			return errorx.New(errorx.CodeBadRequest, errors.New("id is required"))
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid id: %w", err))
		}

		var updates map[string]any
		if err = json.NewDecoder(c.Request().Body).Decode(&updates); err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid json: %w", err))
		}

		for field := range updates {
			if !allowedFields[field] {
				return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid field: %s", field))
			}
		}

		menuItem, err := h.service.Update(c.Request().Context(), id, updates)

		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, menuItem)
	}
}
