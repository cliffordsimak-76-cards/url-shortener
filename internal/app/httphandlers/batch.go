package httphandlers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BatchResponseModel struct {
	CorrelationID string `json:"correlation_id"`
	Short         string `json:"short_url"`
}

// Batch creates a few short URLs by URLs.
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
	return c.JSON(http.StatusCreated, h.ToBatchResponse(createdModels))
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

func (h *HTTPHandler) ToBatchResponse(urlModels []*model.URL) []*BatchResponseModel {
	var result []*BatchResponseModel
	for _, urlModel := range urlModels {
		result = append(result, &BatchResponseModel{
			CorrelationID: urlModel.CorrelationID,
			Short:         h.buildURL(urlModel.Short),
		})
	}
	return result
}
