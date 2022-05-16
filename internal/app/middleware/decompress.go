package middleware

import (
	"compress/gzip"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Decompress() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.Contains(c.Request().Header.Get("Content-Encoding"), "gzip") {
				return next(c)
			}

			gz, err := gzip.NewReader(c.Request().Body)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			defer gz.Close()

			c.Request().Body = gz
			return next(c)
		}
	}
}
