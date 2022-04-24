package httphandlers

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/repository"
	"testing"
)

type testEnv struct {
	urlRepository repository.URLRepository
	httpHandler   *HTTPHandler
}

func newTestEnv(t *testing.T) *testEnv {
	te := &testEnv{}

	te.urlRepository = repository.NewURLRepository()
	te.httpHandler = NewHTTPHandler(
		te.urlRepository,
	)
	return te
}
