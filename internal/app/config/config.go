package config

import (
	"fmt"
	"github.com/caarlos0/env"
	"os"
	"regexp"
	"strings"
)

const PortPattern string = "^([1-9][0-9]{0,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$"

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost:8080"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	if baseURL, ok := os.LookupEnv("BASE_URL"); ok {
		cfg.BaseURL = baseURL
		fmt.Printf("env baseulr: %s\n", baseURL)
	}
	if serverAddress, ok := os.LookupEnv("SERVER_ADDRESS"); ok {
		fmt.Printf("env serverAddress: %s\n", serverAddress)
		cfg.ServerAddress = serverAddress
	}

	err := validatePorts(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func validatePorts(cfg *Config) error {
	substrings := strings.Split(cfg.BaseURL, ":")
	if len(substrings) == 0 {
		return nil
	}
	ok, _ := regexp.MatchString(PortPattern, substrings[2])
	if !ok {
		return fmt.Errorf("base url port has wrong format %s", substrings[2])
	}
	return nil
}
