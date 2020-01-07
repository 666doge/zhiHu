package session


import (
	"sync"
)

type MemorySession struct {
	data map[string]interface{}
	id string
	rwLock sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	ms := &MemorySession{
		id: id,
		data: make(map[string]interface{}, 8),
	}
	return ms
}

func (s *MemorySession) Set (key string, value interface{}) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	s.data[key] = value
	return
}

func (s *MemorySession) Get (key string) (value interface{}, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	value, ok := s.data[key]
	if !ok {
		err = ErrKeyNotExistInSession
		return
	}
	return
}

func (s *MemorySession)Del (key string) (err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	delete(s.data, key)
	return
}

func (s *MemorySession) Save() (err error) {
	return
}