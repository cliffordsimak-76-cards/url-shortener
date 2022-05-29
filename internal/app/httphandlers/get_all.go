package httphandlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type URL struct {
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}

func (h *HTTPHandler) GetAll(c echo.Context) error {
	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlModels, err := h.repository.GetAll(userID)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "error get all")
	}

	if len(urlModels) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	var response []*URL
	for _, url := range urlModels {
		response = append(response, &URL{
			Short:    h.buildURL(url.Short),
			Original: url.Original,
		})
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusOK, response)
}
