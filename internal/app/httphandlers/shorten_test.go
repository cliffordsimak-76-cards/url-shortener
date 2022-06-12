package httphandlers

import (
	"bytes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
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
			name: "success",
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
			req.AddCookie(&http.Cookie{
				Name:  "userID",
				Value: "226d0f8a5fa9180d",
			})
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			h := te.httpHandler.Shorten
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.code, rec.Code)
				compareMessage(t, tt.response, rec.Body.Bytes())
			}
		})
	}
}

func compareMessage(t *testing.T, expected, actual []byte) {
	assert.JSONEq(t, string(expected), string(actual))
}
