package repository

type URLRepository interface {
	Create(id string, url string)
	Get(id string) (string, error)
}
