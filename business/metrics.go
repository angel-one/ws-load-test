package business

import (
	"github.com/angel-one/ws-load-test/models"
	"time"
)

var mainCh chan *models.TestResult
var LatencyTime []time.Time
var Latency []int
var TotalConnection []float64
var Error []float64
var Receive []float64
var ReceiveTime []time.Time
var Send []float64
var SendTime []time.Time
var ErrorTime []time.Time
var TotalConnectionTime []time.Time

var lastErrorTime time.Time
var lastConnectionTime time.Time
var lastSendTime time.Time
var lastLatencyTime time.Time
var lastReceiveTime time.Time

func processEvent(eventTime time.Time, lastTime time.Time, slice []float64)  []float64{
	if len(slice) == 0 {
		slice = append(slice, 0)
	}
	if eventTime.Minute() == lastTime.Minute() {
		var val float64= 0
		if len(slice) > 0 {
			val = slice[len(slice)-1]
		}
		slice = slice[:len(slice)-1]
		slice = append(slice, val+1)
	} else {
		for i := lastTime.Minute() + 1; i < eventTime.Minute()-1; i++ {
			slice = append(slice, 0)
		}
		slice = append(slice, 1)
	}
	return slice
}

func processLatency(eventTime time.Time, lastTime time.Time, slice []int, latency int)  []int{
	if len(slice) == 0 {
		slice = append(slice, 0)
	}
	if eventTime.Minute() == lastTime.Minute() {
		var val int= 0
		if len(slice) > 0 {
			val = slice[len(slice)-1]
		}
		slice = slice[:len(slice)-1]
		slice = append(slice, (val+latency)/2)
	} else {
		for i := lastTime.Minute() + 1; i < eventTime.Minute()-1; i++ {
			slice = append(slice, 0)
		}
		slice = append(slice, latency)
	}
	return slice
}

func Init() {
	lastErrorTime = time.Now()
	lastConnectionTime = time.Now()
	lastSendTime = time.Now()
	lastLatencyTime = time.Now()
	lastReceiveTime = time.Now()
	for {
		data := <-mainCh
		if data.EventType == "start" {
			TotalConnection = processEvent(data.EventTime, lastConnectionTime, TotalConnection)
			lastConnectionTime = data.EventTime
		} else if data.EventType == "done" {
			TotalConnection = processEvent(data.EventTime, lastConnectionTime, TotalConnection)
			lastConnectionTime = data.EventTime
		} else if data.EventType == "error" {
			Error = processEvent(data.EventTime, lastErrorTime, Error)
			lastErrorTime = data.EventTime
		} else if data.EventType == "send" {
			Send = processEvent(data.EventTime, lastSendTime, Send)
			lastSendTime = data.EventTime
			Latency = processLatency(data.EventTime, lastLatencyTime, Latency, data.Latency)
			lastLatencyTime = data.EventTime
		} else if data.EventType == "receive" {
			Receive = processEvent(data.EventTime, lastReceiveTime, Send)
			lastReceiveTime = data.EventTime
			Latency = processLatency(data.EventTime, lastLatencyTime, Latency, data.Latency)
			lastLatencyTime = data.EventTime
		}
	}
}

func HandleMetricsLatency() ([]int, []time.Time) {
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
