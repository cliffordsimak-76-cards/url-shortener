package httphandlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildUrl(t *testing.T) {
	te := newTestEnv(t)
	tests := []struct {
		identifier string
		want       string
	}{
		{
			identifier: "8e43",
			want:       "http://localhost:8080/8e43",
		},
	}
	for _, tCase := range tests {
		resp := te.httpHandler.buildURL(tCase.identifier)
		assert.Equal(t, tCase.want, resp)
	}
}
