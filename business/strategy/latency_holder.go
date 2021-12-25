package strategy

import (
	"sync"
	"time"
)

type LatencyHolder struct {
	sync.RWMutex
	sendTime   time.Time
	isSendDone bool
}

func (m *LatencyHolder) Get() (bool, time.Time) {
	m.Lock()
	a,b := m.isSendDone, m.sendTime
	m.Unlock()
	return a,b
}

func (m *LatencyHolder) GetBool() bool {
	m.Lock()
	a := m.isSendDone
	m.Unlock()
	return a
}

func (m *LatencyHolder) GetTime() time.Time {
	m.Lock()
	a := m.sendTime
	m.Unlock()
	return a
}

func (m *LatencyHolder) Set(a bool, b time.Time) {
	m.Lock()
	m.sendTime = b
	m.isSendDone = a
	m.Unlock()
}

func (m *LatencyHolder) SetBool(a bool) {
	m.Lock()
	m.isSendDone = a
	m.Unlock()
}
