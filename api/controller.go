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
	body := "<h1>Hello Carrot!</h1>"
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
	graph := charts.DrawChart(data, timeSeries, "time", "latency")
	graph.Render(chart.PNG, ctx.Writer)
}
