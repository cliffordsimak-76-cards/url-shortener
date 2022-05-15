package main

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	cfg, err := setConfig()
	if err != nil {
		log.Fatal("error running server: ", err)
	}

	if err = app.Run(cfg); err != nil {
		log.Fatal("error running server", err)
	}
}

func setConfig() (*config.Config, error) {
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

	baseURLPort := strings.Split(cfg.BaseURL, ":")[2]
	ok, _ := regexp.MatchString(cfg.PortPattern, baseURLPort)
	if !ok {
		return nil, fmt.Errorf("base url port has wrong format %s", baseURLPort)
	}

	return cfg, nil
}
