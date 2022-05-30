package repository

import (
	"bufio"
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

type InFile struct {
	cache     map[string]string
	userCache map[string][]*model.URL
	encoder   *json.Encoder
	mutex     *sync.Mutex
}

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

func (s *InFile) Create(
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

func (s *InFile) CreateBatch(urlModels []*model.URL) error {
	panic("implement me")
}

func (s *InFile) Get(
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

func (s *InFile) GetAll(
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
