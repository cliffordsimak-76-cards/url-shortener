package httphandlers

import (
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func (h *HTTPHandler) Post(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if len(body) == 0 {
			return c.String(http.StatusBadRequest, "body is empty")
		}

		URL := string(body)
		err = validateURL(URL)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		urlIdentifier := uuid.New().String()
		err = h.repository.Create(urlIdentifier, URL)
		if errors.Is(err, repository.ErrAlreadyExists) {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err != nil {
			log.Error(err)
			return c.String(http.StatusBadRequest, "error create in db")
		}

		return c.String(http.StatusCreated, makeShortLink(cfg.BaseURL, urlIdentifier))
	}
}
