package httphandlers

import (
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func (h *HTTPHandler) Shorten(c echo.Context) error {
	var request *ShortenRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := validateURL(request.URL)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	URL := adapters.ToModel(userID, request.URL)
	urlID, err := h.create(URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(
		http.StatusCreated,
		adapters.ToShortenResponse(h.buildURL(urlID)),
	)
}
