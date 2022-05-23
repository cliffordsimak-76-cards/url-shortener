package repository

import (
	"sync"
)

type InMemory struct {
	cache map[string]map[string]string
	mutex *sync.Mutex
}

func NewInMemory() Repository {
	return &InMemory{
		cache: make(map[string]map[string]string),
		mutex: &sync.Mutex{},
	}
}

func (s *InMemory) Create(
	userID string,
	id string,
	url string,
) error {
	if _, ok := s.cache[userID][id]; ok {
		return ErrAlreadyExists
	}
	s.mutex.Lock()
	urlByID, ok := s.cache[userID]
	if !ok {
		urlByID = map[string]string{}
		s.cache[userID] = urlByID
	}
	urlByID[id] = url
	s.mutex.Unlock()
	return nil
}

func (s *InMemory) Get(
	userID string,
	id string,
) (string, error) {
	s.mutex.Lock()
	URL, ok := s.cache[userID][id]
	s.mutex.Unlock()
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}

func (s *InMemory) GetAll(
	userID string,
) (map[string]string, error) {
	s.mutex.Lock()
	urlMap, ok := s.cache[userID]
	s.mutex.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return urlMap, nil
}
