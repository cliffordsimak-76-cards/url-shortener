package httphandlers

import (
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h *HTTPHandler) Get(c echo.Context) error {
	URL, err := h.repository.Get(c.Param("id"))
	if err != nil {
		log.Error(err)
		if errors.Is(err, repository.ErrNotFound) {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusBadRequest, "error get")
	}

	c.Response().Header().Set(echo.HeaderLocation, URL)
	return c.NoContent(http.StatusTemporaryRedirect)
}
