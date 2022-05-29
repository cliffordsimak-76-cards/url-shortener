package app

import (
	"database/sql"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/middleware"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func Run(cfg *config.Config) error {
	db, err := sql.Open("postgres", cfg.DatabaseDSN)
	if err != nil {
		return err
	}

	var repo repository.Repository
	if cfg.FileStoragePath != "" {
		repo, err = repository.NewInFile(db, cfg.FileStoragePath)
		if err != nil {
			return err
		}
	} else {
		repo = repository.NewInMemory(db)
	}
	httpHandler := httphandlers.NewHTTPHandler(repo, cfg)

	e := echo.New()
	e.GET("/ping", httpHandler.Ping)
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
