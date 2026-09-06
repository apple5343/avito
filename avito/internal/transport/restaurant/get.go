package restaurant

import (
	"fmt"
	"net/http"
	"strconv"

	"avito/internal/model"
	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid id: %w", err))
		}
		restaurant, err := h.service.GetByID(c.Request().Context(), id)
		if err != nil {
			return err
		}
		restaurant.APIKey = ""
		restaurant.CallbackURL = ""
		return c.JSON(http.StatusOK, restaurant)
	}
}

func (h *Handler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		filter := &model.RestaurantFilter{}
		if name := c.QueryParam("name"); name != "" {
			filter.Name = &name
		}
		if city := c.QueryParam("city"); city != "" {
			filter.City = &city
		}
		if cuisine := c.QueryParam("cuisine"); cuisine != "" {
			filter.Cuisine = &cuisine
		}
		if isActive, err := strconv.ParseBool(c.QueryParam("is_active")); err == nil {
			filter.IsActive = &isActive
		} else if c.QueryParam("is_active") != "" {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid is_active: %w", err))
		}
		restaurants, err := h.service.List(c.Request().Context(), filter)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, restaurants)
	}
}
