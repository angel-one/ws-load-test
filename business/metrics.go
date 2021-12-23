package business

import (
	"github.com/angel-one/ws-load-test/models"
	"time"
)

var mainCh chan *models.TestResult
var LatencyTime []time.Time
var Latency []float64
var Error []float64
var ErrorTime []time.Time
var Receive []float64
var ReceiveTime []time.Time
var Send []float64
var SendTime []time.Time
var TotalConnection []float64
var TotalConnectionTime []time.Time


func Init() {
	for {
		data := <-mainCh
		if data.EventType == "start" {
			TotalConnectionTime = append(TotalConnectionTime, data.EventTime)
			TotalConnection = append(TotalConnection, 1)
		} else if data.EventType == "end" {
			TotalConnectionTime = append(TotalConnectionTime, data.EventTime)
			TotalConnection = append(TotalConnection, -1)
		} else if data.EventType == "error" {
			ErrorTime = append(ErrorTime, data.EventTime)
			Error = append(Error, 1)
		} else if data.EventType == "send" {
			SendTime = append(SendTime, data.EventTime)
			Send = append(Send, 1)
			LatencyTime = append(LatencyTime, data.EventTime)
			Latency = append(Latency, data.Latency)
		} else if data.EventType == "receive" {
			ReceiveTime = append(ReceiveTime, data.EventTime)
			Receive = append(Receive, 1)
			LatencyTime = append(LatencyTime, data.EventTime)
			Latency = append(Latency, data.Latency)
		}
	}
}

func HandleMetricsLatency() ([]float64, []time.Time) {
	return Latency, LatencyTime
}

func HandleMetricsError() ([]float64, []time.Time) {
	return Error, ErrorTime
}

func HandleMetricsSend() ([]float64, []time.Time) {
	return Send, SendTime
}

func HandleMetricsReceive() ([]float64, []time.Time) {
	return Receive, ReceiveTime
}

func HandleMetricsConnection() ([]float64, []time.Time) {
	return TotalConnection, TotalConnectionTime
}

func SetMainChannel(ch chan *models.TestResult) {
	mainCh = ch
}
