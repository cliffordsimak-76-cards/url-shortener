package app

import (
	"database/sql"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/middleware"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Run(cfg *config.Config) error {
	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	repo, err := initRepo(cfg, db)
	if err != nil {
		log.Fatal(err)
		return err
	}

	httpHandler := httphandlers.NewHTTPHandler(cfg, repo, db)

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

func initDB(
	cfg *config.Config,
) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(schema); err != nil {
		return nil, err
	}
	return db, nil
}

func initRepo(
	cfg *config.Config,
	db *sql.DB,
) (repository.Repository, error) {
	if cfg.DatabaseDSN != "" {
		return repository.NewInDatabase(db), nil
	}
	if cfg.FileStoragePath != "" {
		repo, err := repository.NewInFile(cfg.FileStoragePath)
		if err != nil {
			return nil, err
		}
		return repo, nil
	}
	return repository.NewInMemory(), nil
}

var schema = `
	CREATE TABLE IF NOT EXISTS urls (
		id serial primary key,
		base_url text not null unique,
		short_url text not null 
	);
	CREATE TABLE IF NOT EXISTS users_url(
	  user_id text not null,
	  url_id int not null references urls(id),
	  CONSTRAINT unique_url UNIQUE (user_id, url_id)
	);
	`
