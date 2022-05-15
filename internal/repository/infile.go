package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

type InFile struct {
	cache   map[string]string
	encoder *json.Encoder
}

func NewInFile(filePath string) (Repository, error) {
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

	return &InFile{
		cache:   cache,
		encoder: json.NewEncoder(file),
	}, nil
}

func (s *InFile) Create(id string, url string) error {
	if _, ok := s.cache[id]; ok {
		return ErrAlreadyExists
	}
	s.cache[id] = url

	data := make(map[string]string, 1)
	data[id] = url
	return s.encoder.Encode(&data)
}

func (s *InFile) Get(id string) (string, error) {
	URL, ok := s.cache[id]
	if !ok {
		return "", ErrNotFound
	}
	return URL, nil
}
