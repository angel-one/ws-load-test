package business

import (
	"github.com/angel-one/ws-load-test/models"
	"time"
)

var mainCh chan *models.TestResult
var TimeCh []time.Time
var Latency []int64

func Init() {
	go func() {
		data := <-mainCh
		TimeCh = append(TimeCh, data.SendTimeLatest)
		Latency = append(Latency, data.TimeDiff)
	}()
}

func HandleMetricsLatency() ([]int64, []time.Time) {
	return Latency, TimeCh
}

func SetMainChannel(ch chan *models.TestResult) {
	mainCh = ch
}
