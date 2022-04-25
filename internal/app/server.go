package app

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"net/http"
)

const port = ":8080"

var URLsByID = map[string]string{}

func Run() error {
	urlRepository := repository.NewURLRepository()
	httpHandler := httphandlers.NewHTTPHandler(urlRepository)

	http.HandleFunc("/", httpHandler.HandleRequest)
	http.ListenAndServe(port, nil)

	return nil
}
