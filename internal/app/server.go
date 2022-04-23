package app

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"path"
)

const port = ":8080"

var urlsByID = map[string]string{}

func Run() error {
	http.HandleFunc("/", HandleRequest)
	http.ListenAndServe(port, nil)
	return nil
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetHandler(w, r)
	case http.MethodPost:
		PostHandler(w, r)
	default:
		BadReQuestHandler(w, r)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	_, shortUrl := path.Split(r.URL.Path)
	if shortUrl == "" {
		BadReQuestHandler(w, r)
		return
	}

	url, found := urlsByID[shortUrl]
	if !found {
		BadReQuestHandler(w, r)
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(generateShortUrl(string(body))))
}

func generateShortUrl(url string) string {
	if value, found := urlsByID[url]; found {
		return value
	}
	shorUrl := uuid.New().String()
	urlsByID[shorUrl] = url
	return shorUrl
}

func BadReQuestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}
