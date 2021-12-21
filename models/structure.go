package models

import (
	"time"
)

type TestResult struct {
	SendTimeLatest    time.Time
	ReceiveTimeLatest time.Time
	TimeDiff          int64
	ReceivedMsgCount  int64
	SendMsgCount      int64
}
