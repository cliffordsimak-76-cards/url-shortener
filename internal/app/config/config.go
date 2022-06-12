package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env"
)

const (
	UserCookieName string = "userID"
	UserIDLen      int    = 8
	SecretKey      string = "secret key"
)

type Config struct {
	ServerAddress   string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL         string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:""`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	parseFlags(cfg)

	fmt.Println(cfg)
	return cfg, nil
}

func parseFlags(cfg *Config) {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "address to listen on")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base URL for short link")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "database address")
	flag.Parse()
}

func (cfg *Config) String() string {
	return fmt.Sprintf(
		"SERVER_ADDRESS: %s\n"+
			"BASE_URL: %s\n"+
			"FILE_STORAGE_PATH: %s\n"+
			"DATABASE_DSN: %s\n",
		cfg.ServerAddress, cfg.BaseURL, cfg.FileStoragePath, cfg.DatabaseDSN,
	)
}
