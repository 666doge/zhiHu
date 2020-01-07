package session

import (
	"testing"
)

func TestSession(t *testing.T) {
	sm := NewMemorySessionManager()
	ms, _ := sm.CreateSession()
	_ = ms.Set("name", "xushengnan")
	v, _ := ms.Get("name")
	if (v != "xushengnan") {
		t.Error("result is wrong")
	}
	_ = ms.Del("name")
	v, _ = ms.Get("name")
	if (v != nil) {
		t.Error("del failed")
	}
}