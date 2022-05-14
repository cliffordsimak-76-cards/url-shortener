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

		urlIdentifier := uuid.New().String()
		shortURL := host + urlIdentifier
		err = h.urlRepository.Create(urlIdentifier, string(body))
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		return c.String(http.StatusCreated, shortURL)
	}
}
