package models

import (
	"time"
)

type TestResult struct {
	ID                int64
	SendTimeFirst     time.Time
	ReceiveTimeFirst  time.Time
	SendTimeLatest    time.Time
	ReceiveTimeLatest time.Time
	Latency           float64
	ReceivedMsgCount  int64
	SendMsgCount      int64
	hasEnded          bool
}
