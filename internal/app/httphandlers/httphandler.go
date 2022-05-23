package httphandlers

import (
	"errors"
	"fmt"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/gommon/log"
	"net/url"
	"strings"
)

type HTTPHandler struct {
	repository repository.Repository
	cfg        *config.Config
}

func NewHTTPHandler(
	repository repository.Repository,
	cfg *config.Config,
) *HTTPHandler {
	return &HTTPHandler{
		repository: repository,
		cfg:        cfg,
	}
}

func validateURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}
	return nil
}

func (h *HTTPHandler) generateUrlID(URL string) (string, error) {
	urlID := utils.StringToMD5(URL)
	err := h.repository.Create(urlID, URL)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExists) {
			return "", err
		}
		log.Error(err)
		return "", fmt.Errorf("error create in db")
	}
	return urlID, nil
}

func (h *HTTPHandler) buildURL(id string) string {
	return strings.Join([]string{h.cfg.BaseURL, id}, "/")
}
