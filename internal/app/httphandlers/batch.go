package httphandlers

import (
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/url"
)

func (h *HTTPHandler) Batch(c echo.Context) error {
	var request []*adapters.BatchRequestModel
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userID, err := extractUserID(c.Request())
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = validateURLBatch(request)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	urlModels := adapters.ToModels(userID, request)
	createdModels, err := h.createBatch(urlModels)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(http.StatusCreated, adapters.ToBatchResponse(createdModels))
}

func validateURLBatch(request []*adapters.BatchRequestModel) error {
	for _, model := range request {
		_, err := url.ParseRequestURI(model.Original)
		if err != nil {
			return err
		}
	}
	return nil
}
