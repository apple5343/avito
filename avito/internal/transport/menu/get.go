package menu

import (
	"fmt"
	"net/http"
	"strconv"

	"avito/internal/model"
	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetItemByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		menuItem, err := h.service.GetItemByID(c.Request().Context(), id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, menuItem)
	}
}

func (h *Handler) GetMenu() echo.HandlerFunc {
	return func(c echo.Context) error {
		restaurantID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		filter := &model.MenuFilter{}
		if category := c.QueryParam("category"); category != "" {
			filter.Category = &category
		}
		var isAvailable bool
		if isAvailable, err = strconv.ParseBool(c.QueryParam("is_available")); err == nil {
			filter.IsAvailable = &isAvailable
		} else if c.QueryParam("is_available") != "" {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid is_available: %w", err))
		}
		menuItems, err := h.service.GetMenu(c.Request().Context(), restaurantID, filter)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, menuItems)
	}
}
