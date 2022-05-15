package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostURL(t *testing.T) {
	te := newTestEnv(t)
	type want struct {
		code int
	}
	tests := []struct {
		name  string
		value string
		want  want
	}{
		{
			name:  "body is empty",
			value: "",
			want:  want{code: 400},
		},
		{
			name:  "with param",
			value: "https://www.yandex.ru",
			want:  want{code: 201},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.value))
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			cfg := &config.Config{
				BaseURL: "http://localhost:8080",
			}
			h := te.httpHandler.Post(cfg)
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
				require.NotNil(t, rec.Body.String())
			}
		})
	}
}
