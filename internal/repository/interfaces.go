package repository

type UrlRepository interface {
	Create(id string, url string)
	Get(id string) (string, error)
}
