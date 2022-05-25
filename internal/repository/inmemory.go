package repository

import (
	"database/sql"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	_ "github.com/lib/pq"
	"sync"
)

type InMemory struct {
	db        *sql.DB
	cache     map[string]string
	userCache map[string][]*model.URL
	mutex     *sync.Mutex
}

func NewInMemory(db *sql.DB) Repository {
	return &InMemory{
		db:        db,
		cache:     make(map[string]string),
		userCache: make(map[string][]*model.URL),
		mutex:     &sync.Mutex{},
	}
}

func (s *InMemory) Create(
	userID string,
	id string,
	url string,
) error {
	if _, ok := s.cache[id]; ok {
		return ErrAlreadyExists
	}
	s.mutex.Lock()
	s.cache[id] = url
	_, ok := s.userCache[userID]
	if !ok {
		s.userCache[userID] = make([]*model.URL, 0)
	}
	s.userCache[userID] = append(s.userCache[userID], &model.URL{
		Short:    id,
		Original: url,
	})
	s.mutex.Unlock()
	return nil
}

func (s *InMemory) Get(
	id string,
) (string, error) {
	s.mutex.Lock()
	URL, ok := s.cache[id]
	s.mutex.Unlock()
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}

func (s *InMemory) GetAll(
	userID string,
) ([]*model.URL, error) {
	s.mutex.Lock()
	urls, ok := s.userCache[userID]
	s.mutex.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return urls, nil
}

func (s *InMemory) Ping() error {
	if err := s.db.Ping(); err != nil {
		return err
	}
	return nil
}
