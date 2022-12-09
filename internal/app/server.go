package app

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/grpchandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/httphandlers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/middleware"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	pb "github.com/cliffordsimak-76-cards/url-shortener/internal/proto"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

// run.
func Run(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

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

	if cfg.TrustedSubnet != "" {
		_, trustedNet, err := net.ParseCIDR(cfg.TrustedSubnet)
		if err != nil {
			log.Fatal(err)
			return err
		}

		e.POST("/api/internal/stats", httpHandler.GetStats)
		e.Use(middleware.IPFilter(trustedNet))
	}

	go func() {
		fmt.Println(http.ListenAndServe(cfg.PprofAddress, nil))
	}()

	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			log.Fatalf("GRPC server net.Listen: %v", err)
		}

		s := grpc.NewServer()
		pb.RegisterShortenerServer(s, grpchandlers.NewGrpcServer(repo))

		log.Printf("GRPC server started on %v", "3200")

		if err := s.Serve(listen); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		<-signalChan

		log.Print("Shutting down...")

		cancel()
		if err = e.Shutdown(ctx); err != nil && err != ctx.Err() {
			e.Logger.Fatal(err)
		}

		if err = db.Close(); err != nil {
			log.Fatal(err)
		}

		close(deleteTasks)
	}()

	if cfg.EnabledHTTPS {
		if err = utils.CheckCerts(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(e.StartTLS(cfg.ServerAddress, utils.CertFile, utils.KeyFile))
	} else {
		e.Logger.Fatal(e.Start(cfg.ServerAddress))
	}

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
