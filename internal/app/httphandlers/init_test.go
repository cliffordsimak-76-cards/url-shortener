package httphandlers

import (
	"testing"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
)

type testEnv struct {
	inMemoryRepo repository.Repository
	httpHandler  *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	cfg := &config.Config{
		BaseURL: "http://localhost:8080",
	}

	te.inMemoryRepo = repository.NewInMemory()
	te.httpHandler = NewHTTPHandler(
		cfg,
		te.inMemoryRepo,
		nil,
		nil,
	)
	return te
}
