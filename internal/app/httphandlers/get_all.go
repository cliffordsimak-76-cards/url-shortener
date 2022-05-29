package httphandlers

import (
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func (h *HTTPHandler) GetAll(c echo.Context) error {
	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlModels, err := h.repository.GetAll(userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, repository.ErrNotFound) {
			return c.NoContent(http.StatusNoContent)
		}
		return c.String(http.StatusBadRequest, "error get all")
	}
	if len(urlModels) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusOK, adapters.ToGetAllResponse(urlModels))
}
