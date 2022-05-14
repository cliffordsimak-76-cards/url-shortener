package repository

type URLRepository interface {
	Create(id string, url string) error
	Get(id string) (string, error)
}
