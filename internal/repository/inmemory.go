package repository

import "sync"

type InMemory struct {
	cache map[string]string
	mutex *sync.Mutex
}

func NewInMemory() Repository {
	return &InMemory{
		cache: make(map[string]string),
		mutex: &sync.Mutex{},
	}
}

func (s *InMemory) Create(id string, url string) error {
	if _, ok := s.cache[id]; ok {
		return ErrAlreadyExists
	}
	s.mutex.Lock()
	s.cache[id] = url
	s.mutex.Unlock()
	return nil
}

func (s *InMemory) Get(id string) (string, error) {
	s.mutex.Lock()
	URL, ok := s.cache[id]
	s.mutex.Unlock()
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}
