package httphandlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HTTPHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		URL, err := h.urlRepository.Get(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Errorf("error get URL: %s", err).Error())
		}

		c.Response().Header().Set("Location", URL)
		return c.NoContent(http.StatusTemporaryRedirect)
	}
}
