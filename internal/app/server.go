package app

import (
	"io"
	"net/http"
	"path"
	"strconv"
)

const port = ":8080"

var URLsByID = map[int]string{}
var index int

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

	id_input, err := strconv.Atoi(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	URL, found := URLsByID[id_input]
	if !found {
		BadReQuestHandler(w, r)
	}

	w.Header().Set("Location", URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(generateShortURL(string(body)))))
}

func generateShortURL(URL string) int {
	//id_input, err := strconv.Atoi(URL)
	//if err != nil {
	//	//w.WriteHeader(http.StatusBadRequest)
	//	//	return
	//}
	//if value, found := URLsByID[id_input]; found {
	//	return value
	//}
	//shortURL := uuid.New().String()
	//URLsByID[shortURL] = URL

	index = len(URLsByID)
	URLsByID[index] = URL
	return index
}

func BadReQuestHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusBadRequest)
}
