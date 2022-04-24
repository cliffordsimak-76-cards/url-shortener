package httphandlers

import (
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

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.value))
			rec := httptest.NewRecorder()
			h := http.HandlerFunc(te.httpHandler.postURL)
			h.ServeHTTP(rec, req)
			result := rec.Result()

			if result.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, rec.Code)
			}
		})
	}
}
