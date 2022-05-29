package repository

import (
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
)

type Repository interface {
	Create(userID string, id string, url string) error
	Get(id string) (string, error)
	GetAll(userID string) ([]*model.URL, error)
}

var (
	ErrNotFound      = errors.New("URL is not found")
	ErrAlreadyExists = errors.New("URL already exists")
)
