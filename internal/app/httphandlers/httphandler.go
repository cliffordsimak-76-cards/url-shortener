package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

type HTTPHandler struct {
	repository repository.Repository
}

func NewHTTPHandler(repository repository.Repository) *HTTPHandler {
	return &HTTPHandler{
		repository: repository,
	}
}
