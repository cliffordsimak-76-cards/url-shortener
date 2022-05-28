package httphandlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h *HTTPHandler) Ping(c echo.Context) error {
	err := h.repository.Ping()
	if err != nil {
		log.Error("error db ping: ", err)
		return c.String(http.StatusInternalServerError, "error ping")
	}

	return c.NoContent(http.StatusOK)
}
