package adapters

import (
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/google/uuid"
)

// ToModels converts userID and URL to *model.URL.
func ToModel(
	userID string,
	URL string,
) *model.URL {
	return &model.URL{
		CorrelationID: uuid.NewString(),
		UserID:        userID,
		Original:      URL,
		Short:         utils.StringToMD5(URL),
	}
}

// BatchRequestModel is the request to Batch method.
type BatchRequestModel struct {
	CorrelationID string `json:"correlation_id"`
	Original      string `json:"original_url"`
}

// ToModels converts []*BatchRequestModel to []*model.URL.
func ToModels(
	userID string,
	batchRequestModels []*BatchRequestModel,
) []*model.URL {
	var urlModels []*model.URL
	for _, batchModel := range batchRequestModels {
		urlModels = append(urlModels, &model.URL{
			CorrelationID: batchModel.CorrelationID,
			UserID:        userID,
			Original:      batchModel.Original,
			Short:         utils.StringToMD5(batchModel.Original),
		})
	}
	return urlModels
}
