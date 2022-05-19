package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
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

func makeShortLink(baseURL string, shortURL string) string {
	return strings.Join([]string{baseURL, shortURL}, "/")
}
