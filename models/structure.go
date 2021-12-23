package models

import (
	"time"
)

type TestResult struct {
	ID                int64
	StartTime         time.Time
	EndTime           time.Time
	SendTimeFirst     time.Time
	ReceiveTimeFirst  time.Time
	SendTimeLatest    time.Time
	ReceiveTimeLatest time.Time
	EventTime         time.Time
	EventType         string
	Latency           int64
	ReceivedMsgCount  int64
	SendMsgCount      int64
	HasEnded          bool
	HasError          bool
}
