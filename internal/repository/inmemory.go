package repository

type InMemory struct {
	urlRepository map[string]string
}

func NewInMemory() Storage {
	return &InMemory{
		urlRepository: make(map[string]string),
	}
}

func (s *InMemory) Create(id string, url string) error {
	if _, ok := s.urlRepository[id]; ok {
		return ErrAlreadyExists
	}
	s.urlRepository[id] = url
	return nil
}

func (s *InMemory) Get(id string) (string, error) {
	URL, ok := s.urlRepository[id]
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}
