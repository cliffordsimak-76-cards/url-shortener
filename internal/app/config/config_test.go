package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidatePorts(t *testing.T) {
	tt := []struct {
		name    string
		cfg     *Config
		isError bool
	}{
		{
			name: "error port",
			cfg: &Config{
				BaseURL: "http://localhost:fdfs",
			},
			isError: true,
		},
		{
			name: "error max limit",
			cfg: &Config{
				BaseURL: "http://localhost:65536",
			},
			isError: true,
		},
		{
			name: "valid min",
			cfg: &Config{
				BaseURL: "http://localhost:0",
			},
			isError: false,
		},
		{
			name: "valid max",
			cfg: &Config{
				BaseURL: "http://localhost:65535",
			},
			isError: false,
		},
		{
			name: "wrong format",
			cfg: &Config{
				BaseURL: ":8080",
			},
			isError: false,
		},
		{
			name: "valid",
			cfg: &Config{
				BaseURL: "http://localhost:8080",
			},
			isError: false,
		},
	}
	for _, tt := range tt {
		err := validatePorts(tt.cfg)
		if tt.isError {
			assert.Error(t, err)
			return
		}
		assert.NoError(t, err)
	}
}
