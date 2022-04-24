package repository

import (
	"fmt"
)

type DB struct {
	urlRepository map[string]string
}

func NewUrlRepository() *DB {
	return &DB{
		urlRepository: make(map[string]string),
	}
}

func (d *DB) Create(id string, url string) {
	d.urlRepository[id] = url
}

func (d *DB) Get(id string) (string, error) {
	URL, found := d.urlRepository[id]
	if !found {
		return "", fmt.Errorf("URL is not found")
	}
	return URL, nil
}
