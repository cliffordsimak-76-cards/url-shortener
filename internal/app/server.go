package app

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
)

const port = ":8080"

func Run() error {
	urlRepository := repository.NewURLRepository()
	httpHandler := httphandlers.NewHTTPHandler(urlRepository)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post())

	e.Logger.Fatal(e.Start(port))

	return nil
}
