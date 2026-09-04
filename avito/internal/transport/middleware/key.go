package middleware

import (
	"avito/internal/util/ctxutil"

	"github.com/labstack/echo/v4"
)

func ApiKeyMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiKey := c.Request().Header.Get("X-API-KEY")
			if len(apiKey) == 0 {
				return echo.ErrUnauthorized
			}
			ctx := ctxutil.ContextWithApiKey(c.Request().Context(), apiKey)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
