package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"testing"
)

type testEnv struct {
	urlRepository repository.UrlRepository
	httpHandler   *HttpHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	te.urlRepository = repository.NewUrlRepository()
	te.httpHandler = NewHttpHandler(
		te.urlRepository,
	)
	return te
}
