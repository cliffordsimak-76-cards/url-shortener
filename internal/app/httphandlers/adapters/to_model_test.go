package adapters

import (
	"testing"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
)

func BenchmarkToModel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		userID := utils.RandomString(10)
		URL := utils.RandomString(10)
		b.StartTimer()
		ToModel(userID, URL)
	}
}
