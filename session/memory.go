package session


import (
	"sync"
)

type MemorySession struct {
	data map[string]interface{}
	id string
	rwLock sync.RWMutex
	isModify bool
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
	s.isModify = true
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
	s.isModify = true
	return
}

func (s *MemorySession) Save() (err error) {
	return
}

func (s *MemorySession) IsModify() bool {
	return s.isModify
}

func (s *MemorySession) GetId() string {
	return s.id
}