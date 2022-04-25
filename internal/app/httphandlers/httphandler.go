package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

const host = "http://localhost:8080/"

type HTTPHandler struct {
	urlRepository repository.URLRepository
}

func NewHTTPHandler(urlRepository repository.URLRepository) *HTTPHandler {
	return &HTTPHandler{
		urlRepository: urlRepository,
	}
}
