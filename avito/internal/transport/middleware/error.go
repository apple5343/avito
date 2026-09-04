package middleware

import (
	"avito/pkg/errorx"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if nil == err {
				return nil
			}
			fmt.Println(err.Error())
			if c.Response().Committed {
				return err
			}

			httpErr := &echo.HTTPError{}
			if errors.As(err, &httpErr) {
				return c.JSON(httpErr.Code, map[string]string{
					"status":  "error",
					"message": httpErr.Error(),
				})
			}

			if commonErr, ok := err.(*errorx.Error); ok && commonErr != nil {
				if isInternalError(commonErr) {
					return c.JSON(commonErr.Code, map[string]string{
						"status":  "error",
						"message": "internal server error",
					})
				}

				return c.JSON(commonErr.Code, map[string]string{
					"status":  "error",
					"message": commonErr.Error(),
				})
			}

			return c.JSON(500, map[string]string{
				"status":  "error",
				"message": "internal server error",
			})
		}
	}
}

func isInternalError(err *errorx.Error) bool {
	return err.Code/100 == 5
}
