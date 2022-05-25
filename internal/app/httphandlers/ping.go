package httphandlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HTTPHandler) Ping(c echo.Context) error {
	err := h.repository.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "error ping")
	}

	return c.NoContent(http.StatusOK)
}
