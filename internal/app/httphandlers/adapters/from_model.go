package adapters

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
)

type ShortenResponse struct {
	Result string `json:"result"`
}

func ToShortenResponse(result string) *ShortenResponse {
	return &ShortenResponse{
		Result: result,
	}
}

type BatchResponseModel struct {
	CorrelationID string `json:"correlation_id"`
	Short         string `json:"short_url"`
}

func ToBatchResponse(urlModels []*model.URL) []*BatchResponseModel {
	var result []*BatchResponseModel
	for _, urlModel := range urlModels {
		result = append(result, &BatchResponseModel{
			CorrelationID: urlModel.CorrelationID,
			Short:         urlModel.Short,
		})
	}
	return result
}
