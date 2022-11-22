package repository

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

// Inline db.
type InFile struct {
	cache     map[string]string
	userCache map[string][]*model.URL
	encoder   *json.Encoder
	mutex     *sync.Mutex
}

// NewInFile.
func NewInFile(
	filePath string,
) (Repository, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	cache := make(map[string]string)
	if fileInfo, _ := file.Stat(); fileInfo.Size() != 0 {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			err = json.Unmarshal(scanner.Bytes(), &cache)
			if err != nil {
				return nil, err
			}
		}
	}

	log.Info("start file repo")
	return &InFile{
		cache:     cache,
		userCache: make(map[string][]*model.URL),
		encoder:   json.NewEncoder(file),
		mutex:     &sync.Mutex{},
	}, nil
}

// Create.
func (s *InFile) Create(
	_ context.Context,
	urlModel *model.URL,
) error {
	if _, ok := s.cache[urlModel.Short]; ok {
		return ErrAlreadyExists
	}
	s.mutex.Lock()
	s.cache[urlModel.Short] = urlModel.Original
	_, ok := s.userCache[urlModel.UserID]
	if !ok {
		s.userCache[urlModel.UserID] = make([]*model.URL, 0)
	}
	s.userCache[urlModel.UserID] = append(s.userCache[urlModel.UserID], &model.URL{
		Short:    urlModel.Short,
		Original: urlModel.Original,
	})
	s.mutex.Unlock()

	data := make(map[string]string, 1)
	data[urlModel.Short] = urlModel.Original
	return s.encoder.Encode(&data)
}

// CreateBatch.
func (s *InFile) CreateBatch(
	_ context.Context,
	urlModels []*model.URL,
) error {
	panic("implement me")
}

// Get.
func (s *InFile) Get(
	_ context.Context,
	id string,
) (*model.URL, error) {
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

// GetAll.
func (s *InFile) GetAll(
	_ context.Context,
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

// UpdateBatch.
func (s *InFile) UpdateBatch(
	_ context.Context,
	task workers.DeleteTask,
) error {
	panic("implement me")
}
