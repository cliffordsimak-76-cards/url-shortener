package httphandlers

import (
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/converter"
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

		shortURL := converter.StringToMD5(request.URL)
		err := h.urlRepository.Create(shortURL, request.URL)
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusCreated, Response{
			Result: makeResultString(cfg.BaseURL, shortURL),
		})
	}
}

func makeResultString(baseURL string, shortURL string) string {
	return baseURL + "/" + shortURL
}
