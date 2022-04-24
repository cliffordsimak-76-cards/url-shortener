package httphandlers

import (
	"github.com/google/uuid"
	"io"
	"net/http"
)

func (h *HttpHandler) postURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		badRequest(w, err.Error())
		return
	}
	if len(body) == 0 {
		badRequest(w, "")
		return
	}

	urlIdentifier := uuid.New().String()
	shortURL := host + urlIdentifier
	h.urlRepository.Create(urlIdentifier, string(body))

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortURL))
}
