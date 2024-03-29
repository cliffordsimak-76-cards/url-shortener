package httphandlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers/adapters"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// request.
type ShortenRequest struct {
	URL string `json:"url"`
}

// response.
type ShortenResponse struct {
	Result string `json:"result"`
}

// Shorten creates a shorl URL by URL.
func (h *HTTPHandler) Shorten(c echo.Context) error {
	ctx := c.Request().Context()

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

	urlModel := adapters.ToModel(userID, request.URL)
	urlID, err := h.create(ctx, urlModel)
	if err != nil {
		log.Errorf("error create in db: %s", err)
		if errors.Is(err, repository.ErrAlreadyExists) {
			return h.sendResponse(c, http.StatusConflict, urlID.Short)
		}
		return c.String(http.StatusInternalServerError, "error create in db")
	}

	return h.sendResponse(c, http.StatusCreated, urlModel.Short)
}

func (h *HTTPHandler) sendResponse(c echo.Context, code int, str string) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.JSON(
		code,
		&ShortenResponse{
			Result: h.buildURL(str),
		},
	)
}
