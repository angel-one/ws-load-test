package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/business"
	"github.com/angel-one/ws-load-test/utils"
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
	graph := charts.DrawChart(utils.Convert64(data), timeSeries, "time", "latency")
	graph.Render(chart.PNG, ctx.Writer)
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
	graph := charts.DrawChart(data, timeSeries, "time", "send message count")
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
	graph := charts.DrawChart(data, timeSeries, "time", "receive message count")
	graph.Render(chart.PNG, ctx.Writer)
}

// Ws Load Test Dashboard godoc
// @Summary Dashboard
// @Tags Dashboard API V1
// @Description Dashboard
// @Router /dashboard [get]
func dashboardController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	//ctx.Writer.Header().Set("Content-Type", "image/png")
	dataReceive, timeSeries1 := business.HandleMetricsReceive()
	dataSend, timeSeries2 := business.HandleMetricsSend()
	//dataConnection, timeSeries3 := business.HandleMetricsConnection()
	//dataError, timeSeries4 := business.HandleMetricsError()
	//dataLatency, timeSeries5 := business.HandleMetricsLatency()
	graph1 := charts.DrawChart(dataReceive, timeSeries1, "time", "message receive count")
	graph2 := charts.DrawChart(dataSend, timeSeries2, "time", "message sent count")
	buffer1 := bytes.NewBuffer([]byte{})
	buffer2 := bytes.NewBuffer([]byte{})
	_ = graph1.Render(chart.PNG, buffer1)
	_ = graph2.Render(chart.PNG, buffer2)
	var images []string
	image1 := base64.StdEncoding.EncodeToString(buffer1.Bytes())
	image2 := base64.StdEncoding.EncodeToString(buffer2.Bytes())
	images = append(images, image1)
	images = append(images, image2)

	ctx.HTML(http.StatusOK, "dashboard.tmpl", gin.H{"images": images})

}
