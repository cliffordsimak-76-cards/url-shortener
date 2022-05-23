package httphandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HTTPHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := extractUserID(c.Request())
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		URL, err := h.repository.Get(userID, c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		c.Response().Header().Set(echo.HeaderLocation, URL)
		return c.NoContent(http.StatusTemporaryRedirect)
	}
}
