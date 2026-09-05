package order

import (
	"fmt"
	"net/http"
	"strconv"

	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func (h *Handler) UpdateStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Status string `json:"status"`
		}
		orderID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid id: %w", err))
		}
		if err = c.Bind(&req); err != nil {
			return errorx.New(errorx.CodeBadRequest, fmt.Errorf("invalid json: %w", err))
		}
		order, err := h.service.UpdateStatus(c.Request().Context(), orderID, req.Status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, order)
	}
}
