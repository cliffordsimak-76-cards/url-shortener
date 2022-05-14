package httphandlers

import (
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/converter"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	Url string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (h *HTTPHandler) Shorten() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *Request
		if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		shortURL := converter.StringToMD5(request.Url)
		err := h.urlRepository.Create(shortURL, request.Url)
		if err != nil {
			return c.String(http.StatusBadRequest, "error create")
		}

		c.Response().Header().Set("Content-Type", "application/json")
		return c.JSON(http.StatusCreated, Response{
			Result: host + shortURL,
		})
	}
}
