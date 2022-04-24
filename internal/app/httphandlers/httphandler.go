package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"net/http"
)

const host = "http://localhost:8080/"

type HttpHandler struct {
	urlRepository repository.UrlRepository
}

func NewHttpHandler(urlRepository repository.UrlRepository) *HttpHandler {
	return &HttpHandler{
		urlRepository: urlRepository,
	}
}

func (h *HttpHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getURL(w, r)
	case http.MethodPost:
		h.postURL(w, r)
	default:
		badRequest(w, "wrong http verb")
	}
}

func badRequest(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}
