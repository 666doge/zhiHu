package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

type MemorySessionManager struct {
	sessionMap map[string]Session
	rwLock sync.RWMutex
}

func NewMemorySessionManager() SessionManager {
	sr := &MemorySessionManager{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}

func (ms *MemorySessionManager) Init(addr string, options ...string) error {
	return nil
}

func (sm *MemorySessionManager) CreateSession()(session Session, err error) {
	sm.rwLock.Lock()
	defer sm.rwLock.Unlock()

	uuid := uuid.NewV4()
	// if err != nil {
	// 	return
	// }

	sessionId := uuid.String()
	session = NewMemorySession(sessionId)

	sm.sessionMap[sessionId] = session
	return
}

func (sm *MemorySessionManager) Get(sessionId string) (session Session, err error) {
	sm.rwLock.Lock()
	defer sm.rwLock.Unlock()

	session, ok := sm.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExists
		return
	}
	return
}
