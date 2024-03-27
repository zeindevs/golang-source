package main

import (
	"fmt"
	"sync"
)

type StoreProducerFunc func() Storer

type Storer interface {
	Push([]byte) (int, error)
	Get(int) ([]byte, error)
}

type MemoryStorage struct {
	mu    sync.RWMutex
	data  [][]byte
	topic string
}

func NewMemoryStore() *MemoryStorage {
	return &MemoryStorage{
		data: [][]byte{},
	}
}

func (s *MemoryStorage) Push(b []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStorage) Get(offset int) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if offset < 0 {
		return nil, fmt.Errorf("offset cannot be smaller than zero")
	}
	if len(s.data)-1 < offset {
		return nil, fmt.Errorf("offset (%d) to high", offset)
	}

	return s.data[offset], nil
}
