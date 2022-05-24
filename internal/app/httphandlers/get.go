package httphandlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h *HTTPHandler) Get(c echo.Context) error {
	userID, err := extractUserID(c.Request())
	log.Info("get user: %s", userID)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	URL, err := h.repository.Get(userID, c.Param("id"))
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	c.Response().Header().Set(echo.HeaderLocation, URL)
	return c.NoContent(http.StatusTemporaryRedirect)
}
