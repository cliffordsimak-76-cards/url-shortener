package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	gzipScheme = "gzip"
)

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

// Write.
func (w *gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Compress middleware.
func Compress(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !strings.Contains(c.Request().Header.Get(echo.HeaderAcceptEncoding), gzipScheme) {
			return next(c)
		}

		gz, err := gzip.NewWriterLevel(c.Response().Writer, gzip.BestSpeed)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer gz.Close()

		res := c.Response()
		res.Header().Add(echo.HeaderContentEncoding, gzipScheme)
		res.Writer = &gzipWriter{Writer: gz, ResponseWriter: res.Writer}
		return next(c)
	}
}
