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

func welcome(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	body := "<h1>Hello Carrot!</h1>"
	fmt.Fprint(ctx.Writer, body)
}

func latency(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsLatency()
	graph := charts.DrawChart(data, timeSeries, "time", "latency")
	graph.Render(chart.PNG, ctx.Writer)
}
