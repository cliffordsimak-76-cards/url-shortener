package httphandlers

import (
	"fmt"
	"net/http"
	"path"
)

func (h *HTTPHandler) getURL(w http.ResponseWriter, r *http.Request) {
	_, urlIdentifier := path.Split(r.URL.Path)
	if urlIdentifier == "" {
		badRequest(w, "no URL identifier")
		return
	}

	URL, err := h.urlRepository.Get(urlIdentifier)
	if err != nil {
		badRequest(w, fmt.Errorf("error get URL: %s", err).Error())
	}

	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}
