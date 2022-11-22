package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// decompress middleware.
func Decompress(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		if !strings.Contains(req.Header.Get(echo.HeaderContentEncoding), gzipScheme) {
			return next(c)
		}

		gz, err := gzip.NewReader(req.Body)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer gz.Close()

		req.Body = gz
		return next(c)
	}
}
