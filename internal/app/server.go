package app

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"path"
)

const host = "http://localhost"
const port = ":8080"

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
	fmt.Printf("запросили URL: %s\n", r.URL)
	fmt.Printf("url path: %s\n", r.URL.Path)
	_, shortURL := path.Split(r.URL.Path)
	if shortURL == "" {
		BadReQuestHandler(w, r)
		return
	}

	URL, found := URLsByID[shortURL]
	if !found {
		BadReQuestHandler(w, r)
	}

	fmt.Printf("я вернул полный сайт: %s", URL)
	http.Redirect(w, r, URL, http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	fmt.Printf("прислали: %s\n", string(body))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	link := generateShortURL(string(body))
	fmt.Printf("я вернул: %s\n", link)
	w.Write([]byte(link))
}

func generateShortURL(URL string) string {
	if value, found := URLsByID[URL]; found {
		return value
	}
	guid := uuid.New().String()
	shortURL := host + port + "/" + guid
	URLsByID[guid] = URL
	return shortURL
}

func BadReQuestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}
