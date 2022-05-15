package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestBudgetQuantumFromPb(t *testing.T) {
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
		resp := MakeResultString(tCase.baseURL, tCase.identifier)
		assert.Equal(t, tCase.want, resp)
	}
}
