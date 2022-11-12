package repository

import (
	"context"
	"sync"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

// in memroy db.
type InMemory struct {
	cache     map[string]string
	userCache map[string][]*model.URL
	mutex     *sync.Mutex
}

func NewInMemory() Repository {
	log.Info("start memory repo")
	return &InMemory{
		cache:     make(map[string]string),
		userCache: make(map[string][]*model.URL),
		mutex:     &sync.Mutex{},
	}
}

func (s *InMemory) Create(
	url *model.URL,
) error {
	if _, ok := s.cache[url.Short]; ok {
		return ErrAlreadyExists
	}
	s.mutex.Lock()
	s.cache[url.Short] = url.Original
	_, ok := s.userCache[url.UserID]
	if !ok {
		s.userCache[url.UserID] = make([]*model.URL, 0)
	}
	s.userCache[url.UserID] = append(s.userCache[url.UserID], url)
	s.mutex.Unlock()
	return nil
}

func (s *InMemory) CreateBatch(urlModels []*model.URL) error {
	panic("implement me")
}

func (s *InMemory) Get(id string) (*model.URL, error) {
	s.mutex.Lock()
	URL, ok := s.cache[id]
	s.mutex.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return &model.URL{
		Original: URL,
		Short:    id,
	}, nil
}

func (s *InMemory) GetAll(userID string) ([]*model.URL, error) {
	s.mutex.Lock()
	urls, ok := s.userCache[userID]
	s.mutex.Unlock()
	if !ok {
		return nil, ErrNotFound
	}
	return urls, nil
}

func (s *InMemory) UpdateBatch(ctx context.Context, task workers.DeleteTask) error {
	panic("implement me")
}
