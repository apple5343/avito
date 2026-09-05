package order

import (
	"fmt"
	"net/http"
	"strconv"

	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid id: %w", err))
		}
		order, items, err := h.service.GetUserOrder(c.Request().Context(), orderID)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]any{
			"order": order,
			"items": items,
		})
	}
}

func (h *Handler) GetUserOrders() echo.HandlerFunc {
	return func(c echo.Context) error {
		orders, err := h.service.GetUserOrders(c.Request().Context())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, orders)
	}
}
