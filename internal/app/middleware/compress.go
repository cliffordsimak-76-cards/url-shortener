package middleware

import (
	"compress/gzip"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
)

const (
	gzipScheme = "gzip"
)

type gzipWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *gzipWriter) Write(b []byte) (int, error) {
	if w.Header().Get(echo.HeaderContentType) == "" {
		w.Header().Set(echo.HeaderContentType, http.DetectContentType(b))
	}
	return w.Writer.Write(b)
}

func Compress() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !strings.Contains(c.Request().Header.Get(echo.HeaderAcceptEncoding), gzipScheme) {
				return next(c)
			}

			gz, err := gzip.NewWriterLevel(c.Response().Writer, gzip.BestSpeed)
			if err != nil {
				return c.String(http.StatusBadRequest, err.Error())
			}
			defer gz.Close()

			res := c.Response()
			res.Header().Add(echo.HeaderContentEncoding, gzipScheme)
			res.Writer = &gzipWriter{Writer: gz, ResponseWriter: res.Writer}
			return next(c)
		}
	}
}
