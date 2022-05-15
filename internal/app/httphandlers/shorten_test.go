package httphandlers

import (
	"bytes"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestShorten(t *testing.T) {
	te := newTestEnv(t)
	tests := []struct {
		name     string
		request  []byte
		response []byte
		code     int
	}{
		{
			name: "with param",
			request: []byte(`{
								"url":"https://yandex.ru"
							}`),
			response: []byte(`{
								"result":"http://localhost:8080/e9db20b246fb7d3f"
							}`),
			code: 201,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(tt.request))
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			cfg := getConfig()
			h := te.httpHandler.Shorten(cfg)
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.code, rec.Code)
				compareMessage(t, tt.response, rec.Body.Bytes())
			}
		})
	}
}

func getConfig() *config.Config {
	return &config.Config{
		BaseURL:       os.Getenv("BASE_URL"),
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
}

func compareMessage(t *testing.T, expected, actual []byte) {
	assert.JSONEq(t, string(expected), string(actual))
}
