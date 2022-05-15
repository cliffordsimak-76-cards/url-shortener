package repository

type InMemory struct {
	cache map[string]string
}

func NewInMemory() Repository {
	return &InMemory{
		cache: make(map[string]string),
	}
}

func (s *InMemory) Create(id string, url string) error {
	if _, ok := s.cache[id]; ok {
		return ErrAlreadyExists
	}
	s.cache[id] = url
	return nil
}

func (s *InMemory) Get(id string) (string, error) {
	URL, ok := s.cache[id]
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}
