package middleware

import (
	"strconv"

	"avito/internal/util/ctxutil"

	"github.com/labstack/echo/v4"
)

func UserIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userIDStr := c.Request().Header.Get("X-User-Id")
			if len(userIDStr) == 0 {
				return echo.ErrUnauthorized
			}
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				return echo.ErrUnauthorized
			}
			ctx := ctxutil.ContextWithUserID(c.Request().Context(), userID)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
