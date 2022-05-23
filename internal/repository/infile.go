package repository

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"
)

type InFile struct {
	cache   map[string]map[string]string
	encoder *json.Encoder
	mutex   *sync.Mutex
}

func NewInFile(filePath string) (Repository, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	cache := make(map[string]map[string]string)
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
		cache:   cache,
		encoder: json.NewEncoder(file),
		mutex:   &sync.Mutex{},
	}, nil
}

func (s *InFile) Create(
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

	data := make(map[string]map[string]string, 1)
	data[id] = map[string]string{id: url}
	return s.encoder.Encode(&data)
}

func (s *InFile) Get(
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

func (s *InFile) GetAll(
	userID string,
) (map[string]string, error) {
	panic("implement me")
}
