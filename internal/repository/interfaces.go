package repository

import (
	"errors"
)

type Repository interface {
	Create(userID string, id string, url string) error
	Get(userID string, id string) (string, error)
	GetAll(userID string) (map[string]string, error)
}

var (
	ErrNotFound      = errors.New("URL is not found")
	ErrAlreadyExists = errors.New("URL already exists")
)
