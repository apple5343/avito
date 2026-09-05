package menu

import (
	"net/http"

	"avito/internal/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var menuItem model.MenuItem
		if err := c.Bind(&menuItem); err != nil {
			return err
		}
		if err := h.service.Create(c.Request().Context(), &menuItem); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, menuItem)
	}
}
