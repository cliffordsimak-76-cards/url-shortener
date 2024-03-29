package httphandlers

import (
	"errors"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// url.
type URL struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

// GetAll returns all URLs by user.
func (h *HTTPHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlModels, err := h.repository.GetAll(ctx, userID)
	if len(urlModels) == 0 || errors.Is(err, repository.ErrNotFound) {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "error get all")
	}

	response := make([]*URL, 0, len(urlModels))
	for _, url := range urlModels {
		response = append(response, &URL{
			Short:    h.buildURL(url.Short),
			Original: url.Original,
		})
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusOK, response)
}
