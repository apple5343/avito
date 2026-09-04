package restaurant

import (
	"avito/internal/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var restaurant model.Restaurant
		if err := c.Bind(&restaurant); err != nil {
			return err
		}
		if err := h.service.Create(c.Request().Context(), &restaurant); err != nil {
			return err
		}

		return c.JSON(201, restaurant)
	}
}
