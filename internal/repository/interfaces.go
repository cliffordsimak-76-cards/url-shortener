package repository

import "errors"

type Storage interface {
	Create(id string, url string) error
	Get(id string) (string, error)
}

var (
	ErrNotFound      = errors.New("URL is not found")
	ErrAlreadyExists = errors.New("URL already exists")
)
