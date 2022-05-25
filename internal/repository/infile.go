package repository

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	_ "github.com/lib/pq"
	"os"
	"sync"
)

type InFile struct {
	db        *sql.DB
	cache     map[string]string
	userCache map[string][]*model.URL
	encoder   *json.Encoder
	mutex     *sync.Mutex
}

func NewInFile(
	db *sql.DB,
	cfg *config.Config) (Repository, error) {
	file, err := os.OpenFile(cfg.FileStoragePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

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

	return &InFile{
		db:        db,
		cache:     cache,
		userCache: make(map[string][]*model.URL),
		encoder:   json.NewEncoder(file),
		mutex:     &sync.Mutex{},
	}, nil
}

func (s *InFile) Create(
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

	data := make(map[string]string, 1)
	data[id] = url
	return s.encoder.Encode(&data)
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

func (s *InFile) Ping() error {
	if err := s.db.Ping(); err != nil {
		return err
	}
	return nil
}
