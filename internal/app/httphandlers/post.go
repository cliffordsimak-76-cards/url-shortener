package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
)

func (h *HTTPHandler) Post(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	if len(body) == 0 {
		return c.String(http.StatusBadRequest, "body is empty")
	}

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	URL := string(body)
	err = validateURL(URL)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlModel := adapters.ToModel(userID, URL)
	urlID, err := h.create(urlModel)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusCreated, h.buildURL(urlID))
}
