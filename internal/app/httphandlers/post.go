package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

		urlIdentifier := uuid.New().String()
		shortURL := utils.MakeResultString(cfg.BaseURL, urlIdentifier)
		err = h.repository.Create(urlIdentifier, string(body))
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		return c.String(http.StatusCreated, shortURL)
	}
}
