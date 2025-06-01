package store

import (
	"errors"
	"sync"
	"url-shortener/internal/util"
)

type InMemoryStore struct {
	data        map[string]string
	reverseData map[string]string
	mu          sync.RWMutex
}

func NewInMemoryStore(data map[string]string, reverseData map[string]string) *InMemoryStore {
	return &InMemoryStore{data: data, reverseData: reverseData}
}

func (s *InMemoryStore) Save(url string) (code string, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	code, exists := s.reverseData[url]
	if exists {
		return
	}
	code = util.GenerateShortCode()
	s.data[code] = url
	s.reverseData[url] = code
	return
}

func (s *InMemoryStore) Get(code string) (url string, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.data[code]
	if !exists {
		err = errors.New("Not Found")
		return
	}
	return
}
