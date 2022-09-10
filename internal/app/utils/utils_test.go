package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToMD5(t *testing.T) {
	tCases := []struct {
		str  string
		want string
	}{
		{
			str:  "https://yandex.ru",
			want: "e9db20b246fb7d3f",
		},
	}
	for _, tCase := range tCases {
		resp := StringToMD5(tCase.str)
		assert.Equal(t, tCase.want, resp)
	}
}
