package httphandlers

import (
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

		userID, err := extractUserID(c.Request())
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		URL := string(body)
		err = validateURL(URL)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		urlID, err := h.generateUrlID(userID, URL)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusCreated, h.buildURL(urlID))
	}
}
