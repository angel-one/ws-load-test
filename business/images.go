package business

import (
	"bytes"
	"github.com/angel-one/ws-load-test/utils"
	"github.com/angel-one/ws-load-test/utils/charts"
	"github.com/wcharczuk/go-chart"
)

func GetImages() map[string]bytes.Buffer{
	var data  = make(map[string]bytes.Buffer)

	receive, timeSeries := HandleMetricsReceive()
	receiveBuffer := bytes.NewBuffer([]byte{})
	graph := charts.DrawChart(receive, timeSeries, "time", "receive")
	graph.Render(chart.PNG, receiveBuffer)
	data["receive"] = *receiveBuffer

	send, timeSeries := HandleMetricsSend()
	sendBuffer := bytes.NewBuffer([]byte{})
	graph = charts.DrawChart(send, timeSeries, "time", "send")
	graph.Render(chart.PNG, sendBuffer)
	data["send"] = *sendBuffer

	connection, timeSeries := HandleMetricsConnection()
	connectionBuffer := bytes.NewBuffer([]byte{})
	graph = charts.DrawChart(connection, timeSeries, "time", "connection")
	graph.Render(chart.PNG, connectionBuffer)
	data["connections"] = *connectionBuffer

	e, timeSeries := HandleMetricsError()
	errorBuffer := bytes.NewBuffer([]byte{})
	graph = charts.DrawChart(e, timeSeries, "time", "error")
	graph.Render(chart.PNG, errorBuffer)
	data["errors"] = *errorBuffer

	latency, timeSeries := HandleMetricsLatency()
	latencyBuffer := bytes.NewBuffer([]byte{})
	graph = charts.DrawChart(utils.Convert64(latency), timeSeries, "time", "latency")
	graph.Render(chart.PNG, latencyBuffer)
	data["latency"] = *latencyBuffer

	return data
}
