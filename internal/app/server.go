package app

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"path"
)

const port = ":8080"
const host = "http://localhost:8080/"

var URLsByID = map[string]string{}

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
	_, shortURL := path.Split(r.URL.Path)
	if shortURL == "" {
		BadReQuestHandler(w, r)
		return
	}

	URL, found := URLsByID[shortURL]
	if !found {
		BadReQuestHandler(w, r)
	}

	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(generateShortURL(string(body))))
}

func generateShortURL(URL string) string {
	if value, found := URLsByID[URL]; found {
		return value
	}
	urlIdentifier := uuid.New().String()
	URLsByID[urlIdentifier] = URL
	return host + urlIdentifier
}

func BadReQuestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}
