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

func TestRandomString(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "good payload #1",
			args: args{10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(RandomString(tt.args.len)); got != tt.want {
				t.Errorf("RandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(10)
	}
}

func TestRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Good payload",
			args: args{
				min: 50,
				max: 100,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomInt(tt.args.min, tt.args.max); got > tt.args.max || got < tt.args.min {
				t.Errorf("RandomInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
