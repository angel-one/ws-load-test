package models

import (
	"time"
)

type Base struct {
	URL       string
	Proto     string
	Count     int
	Msg       []byte
	Delay     int
	TickDelay int
	Path      string
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration
	ReceivedMsg string
}
