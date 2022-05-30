package adapters

type ShortenResponse struct {
	Result string `json:"result"`
}

func ToShortenResponse(result string) *ShortenResponse {
	return &ShortenResponse{
		Result: result,
	}
}
