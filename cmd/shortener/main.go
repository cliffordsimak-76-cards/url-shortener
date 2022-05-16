package main

import (
	"flag"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	parseFlags(cfg)

	if err = app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}

func parseFlags(cfg *config.Config) {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "address to listen on")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base URL for short link")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
	flag.Parse()
}
