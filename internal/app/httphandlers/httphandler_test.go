package httphandlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeShortLink(t *testing.T) {
	tCases := []struct {
		baseURL    string
		identifier string
		want       string
	}{
		{
			baseURL:    "http://localhost:8080",
			identifier: "8e43",
			want:       "http://localhost:8080/8e43",
		},
	}
	for _, tCase := range tCases {
		resp := makeShortLink(tCase.baseURL, tCase.identifier)
		assert.Equal(t, tCase.want, resp)
	}
}
