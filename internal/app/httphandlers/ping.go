package httphandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Ping returns service state.
func (h *HTTPHandler) Ping(c echo.Context) error {
	if h.db == nil {
		return c.String(http.StatusInternalServerError, "error ping")
	}
	if err := h.db.Ping(); err != nil {
		log.Error("error db ping: ", err)
		return c.String(http.StatusInternalServerError, "error ping")
	}
	return c.NoContent(http.StatusOK)
}
