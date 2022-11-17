package httphandlers

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/gommon/log"
)

// HTTPHandler.
type HTTPHandler struct {
	cfg        *config.Config
	repository repository.Repository
	db         *sql.DB
	deleteCh   chan workers.DeleteTask
}

// NewHTTPHandler.
func NewHTTPHandler(
	cfg *config.Config,
	repository repository.Repository,
	db *sql.DB,
	deleteCh chan workers.DeleteTask,
) *HTTPHandler {
	return &HTTPHandler{
		cfg:        cfg,
		repository: repository,
		db:         db,
		deleteCh:   deleteCh,
	}
}

func (h *HTTPHandler) create(
	ctx context.Context,
	urlModel *model.URL,
) (*model.URL, error) {
	err := h.repository.Create(ctx, urlModel)
	if errors.Is(err, repository.ErrAlreadyExists) {
		urlModel, getErr := h.repository.Get(ctx, urlModel.Short)
		if getErr != nil {
			return nil, fmt.Errorf("error get in db")
		}
		return urlModel, err
	}
	if err != nil {
		return nil, fmt.Errorf("error create in db")
	}
	return urlModel, nil
}

func (h *HTTPHandler) createBatch(
	ctx context.Context,
	urlModels []*model.URL,
) ([]*model.URL, error) {
	err := h.repository.CreateBatch(ctx, urlModels)
	if err != nil {
		log.Error(err)
		if errors.Is(err, repository.ErrAlreadyExists) {
			return nil, err
		}
		return nil, fmt.Errorf("error create in db")
	}
	return urlModels, nil
}

func (h *HTTPHandler) buildURL(id string) string {
	return strings.Join([]string{h.cfg.BaseURL, id}, "/")
}

func validateURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}
	return nil
}

func extractUserID(req *http.Request) (string, error) {
	cookie, err := req.Cookie(config.UserCookieName)
	if err != nil {
		return "", fmt.Errorf("error read cookie")
	}
	data, err := hex.DecodeString(cookie.Value)
	if err != nil {
		log.Errorf("error decode cookie: %s", err)
		return "", fmt.Errorf("error decode cookie")
	}
	return hex.EncodeToString(data[:8]), nil
}
