package httphandlers

import (
	"errors"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Get returns URL by ID.
func (h *HTTPHandler) Get(c echo.Context) error {
	url, err := h.repository.Get(c.Param("id"))
	if err != nil {
		log.Errorf("error get: %s", err)
		if errors.Is(err, repository.ErrNotFound) {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusInternalServerError, "error get")
	}

	if url.Deleted {
		return c.NoContent(http.StatusGone)
	} else {
		c.Response().Header().Set(echo.HeaderLocation, url.Original)
		return c.NoContent(http.StatusTemporaryRedirect)
	}
}
