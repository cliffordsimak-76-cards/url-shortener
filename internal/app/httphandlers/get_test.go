package httphandlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			req.AddCookie(&http.Cookie{
				Name:  "userID",
				Value: "226d0f8a5fa9180d",
			})
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetPath("/:id")
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.value)

			te.inMemoryRepo.Create(&model.URL{
				UserID:   "226d0f8a5fa9180d",
				Original: "https://www.yandex.ru",
				Short:    "a506e095-b901-47db-8b8f-b23f9b1b9e1b",
			})

			h := te.httpHandler.Get
			if assert.NoError(t, h(ctx)) {
				require.Equal(t, tt.want.code, rec.Code)
			}
		})
	}
}
