package middleware

import (
	"compress/gzip"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func Compress() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.Contains(c.Request().Header.Get("Accept-Encoding"), "gzip") {
				return next(c)
			}

			gz, err := gzip.NewWriterLevel(c.Response().Writer, gzip.BestSpeed)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			defer gz.Close()

			c.Response().Writer = gzipWriter{ResponseWriter: c.Response().Writer, Writer: gz}
			return next(c)
		}
	}
}
