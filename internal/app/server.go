package app

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/middleware"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
)

func Run(cfg *config.Config) error {
	var repo repository.Repository
	var err error

	if cfg.FileStoragePath != "" {
		repo, err = repository.NewInFile(cfg.FileStoragePath)
		if err != nil {
			return err
		}
	} else {
		repo = repository.NewInMemory()
	}
	httpHandler := httphandlers.NewHTTPHandler(repo, cfg)

	e := echo.New()
	e.GET("/:id", httpHandler.Get)
	e.GET("/api/user/urls", httpHandler.GetAll)
	e.POST("/", httpHandler.Post)
	e.POST("/api/shorten", httpHandler.Shorten)
	e.Use(middleware.Decompress)
	e.Use(middleware.Compress)
	e.Use(middleware.Cookie)

	e.Logger.Fatal(e.Start(cfg.ServerAddress))

	return nil
}
