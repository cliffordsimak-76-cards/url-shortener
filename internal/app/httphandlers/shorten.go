package httphandlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	Result string `json:"result"`
}

func (h *HTTPHandler) Shorten(c echo.Context) error {
	var request *Request
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = validateURL(request.URL)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlID, err := h.generateURLID(userID, request.URL)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusCreated, Response{
		Result: h.buildURL(urlID),
	})
}
