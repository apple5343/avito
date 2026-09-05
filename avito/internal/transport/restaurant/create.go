package restaurant

import (
	"fmt"
	"net/http"

	"avito/internal/model"
	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var restaurant model.Restaurant
		if err := c.Bind(&restaurant); err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid json: %w", err))
		}
		if err := h.service.Create(c.Request().Context(), &restaurant); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, restaurant)
	}
}
