package utils

import (
	"fmt"
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

func ExampleStringToMD5() {
	out := StringToMD5("https://yandex.ru")
	fmt.Println(out)

	// Output:
	// e9db20b246fb7d3f
}
