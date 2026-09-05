package middleware

import (
	"errors"
	"log"
	"net/http"

	"avito/pkg/errorx"

	"github.com/labstack/echo/v4"
)

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		const statusField = "status"
		const messageField = "message"
		const errorStandard = "error"
		return func(c echo.Context) error {
			err := next(c)
			if nil == err {
				return nil
			}
			log.Println(err.Error())
			if c.Response().Committed {
				return err
			}

			httpErr := &echo.HTTPError{}
			if errors.As(err, &httpErr) {
				return c.JSON(httpErr.Code, map[string]string{
					statusField:  errorStandard,
					messageField: httpErr.Error(),
				})
			}

			commonErr := &errorx.Error{}
			if errors.As(err, &commonErr) {
				if isInternalError(commonErr) {
					return c.JSON(commonErr.Code, map[string]string{
						statusField:  errorStandard,
						messageField: "internal server error",
					})
				}

				return c.JSON(commonErr.Code, map[string]string{
					statusField:  errorStandard,
					messageField: commonErr.Error(),
				})
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{
				statusField:  errorStandard,
				messageField: "internal server error",
			})
		}
	}
}

func isInternalError(err *errorx.Error) bool {
	const internalErrorPrefix = 5
	return err.Code/100 == internalErrorPrefix
}
