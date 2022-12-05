package httphandlers

import (
	"errors"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Stats
func (h *HTTPHandler) GetStats(c echo.Context) error {
	ctx := c.Request().Context()

	stats, err := h.repository.GetStats(ctx)
	if errors.Is(err, repository.ErrNotFound) {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "error get all")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusOK, stats)
}
