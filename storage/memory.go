package storage

import (
	"sync"

	"asset-service/model"
)

type MemoryStorage struct {
	data map[string]model.Asset

	mu sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {

	return &MemoryStorage{
		data: make(map[string]model.Asset),
	}
}

func (m *MemoryStorage) Create(asset model.Asset) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[asset.ID] = asset

	return nil
}

func (m *MemoryStorage) BatchCreate(
	assets []model.Asset,
) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	for _, a := range assets {

		m.data[a.ID] = a
	}

	return nil
}

func (m *MemoryStorage) Delete(id string) bool {

	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.data[id]

	if ok {
		delete(m.data, id)
	}

	return ok
}

func (m *MemoryStorage) GetAll() []model.Asset {

	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]model.Asset, 0)

	for _, a := range m.data {

		result = append(result, a)
	}

	return result
}

func (m *MemoryStorage) Count() int {

	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.data)
}
