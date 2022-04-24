package httphandlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetURL(t *testing.T) {
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
			name:  "without param",
			value: "/",
			want:  want{code: 400},
		},
		{
			name:  "with wrong param",
			value: "/123",
			want:  want{code: 400},
		},
		{
			name:  "with param",
			value: "/a506e095-b901-47db-8b8f-b23f9b1b9e1b",
			want:  want{code: 307},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, tt.value, nil)
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(te.httpHandler.getURL)
			h.ServeHTTP(rec, req)
			result := rec.Result()
			defer result.Body.Close()

			te.urlRepository.Create("a506e095-b901-47db-8b8f-b23f9b1b9e1b", "https://www.yandex.ru")

			assert.Equal(t, tt.want.code, result.StatusCode)
		})
	}
}
