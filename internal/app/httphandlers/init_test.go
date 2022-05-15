package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"testing"
)

type testEnv struct {
	inMemoryRepo repository.Repository
	httpHandler  *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	te.inMemoryRepo = repository.NewInMemory()
	te.httpHandler = NewHTTPHandler(
		te.inMemoryRepo,
	)
	return te
}
