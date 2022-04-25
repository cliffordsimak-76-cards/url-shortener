package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"net/http"
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

func (h *HTTPHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
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
