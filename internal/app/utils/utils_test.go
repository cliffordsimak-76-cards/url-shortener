package utils

import (
	"fmt"
	"math/rand"
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

func BenchmarkStringToMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		url := "https://yandex.ru"
		b.StartTimer()
		StringToMD5(url)
	}
}

func BenchmarkGenerateRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		b.StartTimer()
		GenerateRandom(rand.Intn(100))
	}
}
