package main

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"log"
	"os"
)

func main() {
	cfg := &config.Config{}
	cfg.BaseURL = os.Getenv("BASE_URL")
	cfg.ServerAddress = os.Getenv("SERVER_ADDRESS")

	if err := app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}
