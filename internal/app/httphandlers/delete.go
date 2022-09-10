package httphandlers

import (
	"encoding/json"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) Delete(c echo.Context) error {
	var urlsID []string
	if err := json.NewDecoder(c.Request().Body).Decode(&urlsID); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	h.deleteCh <- workers.DeleteTask{
		UrlsID: urlsID,
		UserID: userID,
	}

	return c.NoContent(http.StatusAccepted)
}
