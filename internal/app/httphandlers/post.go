package httphandlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (h *HTTPHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if len(body) == 0 {
			return c.String(http.StatusBadRequest, "body is empty")
		}

		urlID := uuid.New().String()
		shortURL := host + urlID
		h.urlRepository.Create(urlID, string(body))

		return c.String(http.StatusCreated, shortURL)
	}
}
