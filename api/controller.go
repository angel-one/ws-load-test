package api

import (
	"fmt"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/business"
	"github.com/angel-one/ws-load-test/utils/charts"
	"github.com/gin-gonic/gin"
	"github.com/wcharczuk/go-chart"
	"net/http"
)

// Ws Load Test Welcome godoc
// @Summary Welcome
// @Tags Welcome API V1
// @Description welcome
// @Router /welcome [get]
func welcomeController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	body := "Hello WS Load Test!"
	fmt.Fprint(ctx.Writer, body)
}

// Ws Load Test Latency godoc
// @Summary Latency
// @Tags Latency API V1
// @Description latency
// @Router /latency [get]
func latencyController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsLatency()
	graph := charts.DrawChart(convert64(data), timeSeries, "time", "latency")
	graph.Render(chart.PNG, ctx.Writer)
}

func convert64(ar []int64) []float64 {
	newar := make([]float64, len(ar))
	var v int64
	var i int
	for i, v = range ar {
		newar[i] = float64(v)
	}
	return newar
}


// Ws Load Test Connection godoc
// @Summary Connection
// @Tags Connection API V1
// @Description Connection
// @Router /connection [get]
func connectionController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsConnection()
	graph := charts.DrawChart(data, timeSeries, "time", "total connections")
	graph.Render(chart.PNG, ctx.Writer)
}

// Ws Load Test Error godoc
// @Summary Error
// @Tags Error API V1
// @Description Error
// @Router /error [get]
func errorController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsError()
	graph := charts.DrawChart(data, timeSeries, "time", "errors")
	graph.Render(chart.PNG, ctx.Writer)
}

// Ws Load Test Send godoc
// @Summary Send
// @Tags Send API V1
// @Description Error
// @Router /send [get]
func sendController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsSend()
	graph := charts.DrawChart(data, timeSeries, "time", "errors")
	graph.Render(chart.PNG, ctx.Writer)
}

// Ws Load Test Receive godoc
// @Summary Receive
// @Tags Receive API V1
// @Description Receive
// @Router /receive [get]
	func receiveController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsReceive()
	graph := charts.DrawChart(data, timeSeries, "time", "errors")
	graph.Render(chart.PNG, ctx.Writer)
}
