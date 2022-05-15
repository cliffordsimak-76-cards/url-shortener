package httphandlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	te := newTestEnv(t)
	type want struct {
		code int
		body *string
	}
	tests := []struct {
		name  string
		value string
		want  want
	}{
		{
			name:  "without param",
			value: "",
			want:  want{code: 400},
		},
		{
			name:  "with wrong param",
			value: "123",
			want:  want{code: 400},
		},
		{
			name:  "with param",
			value: "a506e095-b901-47db-8b8f-b23f9b1b9e1b",
			want:  want{code: 307},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.value)

			te.inMemoryRepo.Create("a506e095-b901-47db-8b8f-b23f9b1b9e1b", "https://www.yandex.ru")

			h := te.httpHandler.Get()
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
			}
		})
	}
}
