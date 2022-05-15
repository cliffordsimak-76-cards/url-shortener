package main

import (
	"github.com/caarlos0/env"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"log"
	"os"
)

func main() {
	cfg := setConfig()
	if err := app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}

func setConfig() *config.Config {
	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	if baseURL, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.BaseURL = baseURL
	}

	if serverAddress, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		cfg.ServerAddress = serverAddress
	}

	return cfg
}
