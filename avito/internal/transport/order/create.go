package order

import (
	"fmt"
	"net/http"

	"avito/internal/model"
	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req model.OrderRequest
		if err := c.Bind(&req); err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid json: %w", err))
		}
		order, items, err := h.service.Create(c.Request().Context(), &req)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"order": order,
			"items": items,
		})
	}
}
