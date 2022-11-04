package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/middleware"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
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

	ctx := context.Background()
	deleteTasks := make(chan workers.DeleteTask, 5)
	delService := workers.New(repo, deleteTasks)
	go delService.Run(ctx)

	httpHandler := httphandlers.NewHTTPHandler(cfg, repo, db, deleteTasks)

	e := echo.New()
	e.GET("/ping", httpHandler.Ping)
	e.GET("/:id", httpHandler.Get)
	e.GET("/api/user/urls", httpHandler.GetAll)
	e.DELETE("/api/user/urls", httpHandler.Delete)
	e.POST("/", httpHandler.Post)
	e.POST("/api/shorten", httpHandler.Shorten)
	e.POST("/api/shorten/batch", httpHandler.Batch)
	e.Use(middleware.Cookie)
	e.Use(middleware.Decompress)
	e.Use(middleware.Compress)

	go func() {
		fmt.Println(http.ListenAndServe(cfg.PprofAddress, nil))
	}()

	e.Logger.Fatal(e.Start(cfg.ServerAddress))

	return nil
}

func initDB(
	cfg *config.Config,
) (*sql.DB, error) {
	if cfg.DatabaseDSN == "" {
		return nil, nil
	}
	db, err := sql.Open("pgx", cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(repository.CreateTableQuery); err != nil {
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
		return repository.NewInFile(cfg.FileStoragePath)
	}
	return repository.NewInMemory(), nil
}
