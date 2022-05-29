package repository

import (
	"database/sql"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"sync"
)

type InDatabase struct {
	db    *sql.DB
	mutex *sync.Mutex
}

func (s InDatabase) Create(userID string, id string, url string) error {
	panic("implement me")
}

func (s InDatabase) Get(id string) (string, error) {
	panic("implement me")
}

func (s InDatabase) GetAll(userID string) ([]*model.URL, error) {
	panic("implement me")
}

func NewInDatabase(db *sql.DB) Repository {
	return &InDatabase{
		db:    db,
		mutex: &sync.Mutex{},
	}
}
