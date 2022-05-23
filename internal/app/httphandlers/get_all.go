package httphandlers

import (
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Url struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func (h *HTTPHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := extractUserID(c.Request())
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		urlByID, err := h.repository.GetAll(userID)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				return c.NoContent(http.StatusNoContent)
			}
			return c.String(http.StatusBadRequest, err.Error())
		}

		var response []*Url
		for ID, url := range urlByID {
			response = append(response, &Url{
				Short:    h.buildURL(ID),
				Original: url,
			})
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.JSON(http.StatusOK, response)
	}
}