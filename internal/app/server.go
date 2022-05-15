package app

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) error {
	urlRepository := repository.NewInMemory()
	httpHandler := httphandlers.NewHTTPHandler(urlRepository)

	e := echo.New()
	e.GET("/:id", httpHandler.Get())
	e.POST("/", httpHandler.Post(cfg))
	e.POST("/api/shorten", httpHandler.Shorten(cfg))

	e.Logger.Fatal(e.Start(cfg.ServerAddress))

	return nil
}
