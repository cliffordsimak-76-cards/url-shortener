package config

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env"
)

// Consts.
const (
	// UserCookieName.
	UserCookieName string = "userID"
	// UserIDLen.
	UserIDLen int = 8
	// SecretKey.
	SecretKey string = "secret key"
)

// Config.
type Config struct {
	// ServerAddress.
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	// BaseURL.
	BaseURL string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	// FileStoragePath.
	FileStoragePath string `env:"FILE_STORAGE_PATH" envDefault:""`
	// DatabaseDSN.
	DatabaseDSN string `env:"DATABASE_DSN"`
	// PprofAddress.
	PprofAddress string `env:"PPROF_ADDRESS" envDefault:":6060"`
	// Строковое представление бесклассовой адресации (CIDR).
	TrustedSubnet string `json:"trusted_subnet,omitempty" env:"TRUSTED_SUBNET"`
}

// NewConfig loads 'env' values from environment variables
// and returns Config struct.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to retrieve env variables: %w", err)
	}

	parseFlags(cfg)

	fmt.Println(cfg)
	return cfg, nil
}

// ParseFlags parses the command-line flags from os.Args[1:].
func parseFlags(cfg *Config) {
	flag.StringVar(&cfg.ServerAddress, "a", cfg.ServerAddress, "address to listen on")
	flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base URL for short link")
	flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
	flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "database address")
	flag.StringVar(&cfg.TrustedSubnet, "t", cfg.TrustedSubnet, "sets trusted subnet for incoming requests")
	flag.Parse()
}

// ParseConfigFile parsed config.json file and merger with this config.
func (cfg *Config) ParseConfigFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var configJSON Config
	err = json.Unmarshal(data, &configJSON)
	if err != nil {
		return err
	}

	fmt.Printf("config: %#v\n", configJSON)

	if cfg.DatabaseDSN == "" && configJSON.DatabaseDSN != "" {
		cfg.DatabaseDSN = configJSON.DatabaseDSN
	}
	if cfg.ServerAddress == "" && configJSON.ServerAddress != "" {
		cfg.ServerAddress = configJSON.ServerAddress
	}
	if cfg.FileStoragePath == "" && configJSON.FileStoragePath != "" {
		cfg.FileStoragePath = configJSON.FileStoragePath
	}
	if !cfg.EnabledHTTPS && configJSON.EnabledHTTPS {
		cfg.EnabledHTTPS = true
	}
	if cfg.ServerAddress == "" && configJSON.ServerAddress != "" {
		cfg.ServerAddress = configJSON.ServerAddress
	}
	if cfg.BaseURL == "" && configJSON.BaseURL != "" {
		cfg.BaseURL = configJSON.BaseURL
	}
	if cfg.TrustedSubnet == "" && configJSON.TrustedSubnet != "" {
		cfg.TrustedSubnet = configJSON.TrustedSubnet
	}

	return nil
}

// String returns Config values.
func (cfg *Config) String() string {
	return fmt.Sprintf(
		"SERVER_ADDRESS: %s\n"+
			"BASE_URL: %s\n"+
			"FILE_STORAGE_PATH: %s\n"+
			"DATABASE_DSN: %s\n"+
			"PPROF_ADDRESS: %s\n",
		cfg.ServerAddress,
		cfg.BaseURL,
		cfg.FileStoragePath,
		cfg.DatabaseDSN,
		cfg.PprofAddress,
	)
}
