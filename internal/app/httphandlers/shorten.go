package httphandlers

import (
	"encoding/json"
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (h *HTTPHandler) Shorten(cfg *config.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *Request
		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		urlIdentifier := utils.StringToMD5(request.URL)
		err := h.repository.Create(urlIdentifier, request.URL)
		if errors.Is(err, repository.ErrAlreadyExists) {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusCreated, Response{
			Result: utils.MakeResultString(cfg.BaseURL, urlIdentifier),
		})
	}
}
