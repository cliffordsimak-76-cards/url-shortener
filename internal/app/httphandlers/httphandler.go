package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

type HTTPHandler struct {
	urlRepository repository.Storage
}

func NewHTTPHandler(urlRepository repository.Storage) *HTTPHandler {
	return &HTTPHandler{
		urlRepository: urlRepository,
	}
}
