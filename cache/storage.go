package cache

import (
	"fmt"
	"sync"
	"time"
)

type StorageItem struct {
	data      []byte
	ttl       time.Duration
	createdAt time.Time
}

type Storage struct {
	mu   sync.RWMutex
	data map[string]StorageItem
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]StorageItem),
	}
}

func (s *Storage) Set(key string, value []byte, ttl time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = StorageItem{
		data:      value,
		ttl:       ttl,
		createdAt: time.Now(),
	}
	return nil
}

func (s *Storage) Get(key string) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, exist := s.data[key]
	if !exist {
		return nil, fmt.Errorf("%s not found", key)
	}
	return item.data, nil
}
func (s *Storage) Exist(key string) error {
	_, exist := s.data[key]
	if !exist {
		return fmt.Errorf("%s not found", key)
	}
	return nil
}

func (s *Storage) Del(key string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.Exist(key); err != nil {
		return err
	}
	delete(s.data, key)
	return nil
}

func (s *Storage) TTLCheck() error {
	t := time.NewTicker(time.Millisecond * 100)
	for range t.C {
		for key, val := range s.data {
			if time.Since(val.createdAt) > val.ttl {
				s.mu.Lock()
				delete(s.data, key)
				s.mu.Unlock()
			}
		}
	}
	defer t.Stop()
	return nil
}
