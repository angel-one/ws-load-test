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

// PositionReport godoc
// @Summary Position Report
// @Tags Position Report API V1
// @Description Provides position, trade, order and other reports
// @Param request body string false "201|A599059|A599059|26F638C8EADE24F38A6480D7EB53DA||||"
// @Success 200 {string} Token "251|0|0000306|nse_cm|1594|A599059|CNC|INFY-EQ|INFY||EQ||-1||1|2|-1|INFOSYS LIMITED|1|5|1|1|1|1|1|0|171555|0|0|0|0|0|171555|0|171555|-171555|1|171555|0|0|0|171555|0|171555^nse_cm|22|A599059|CNC|ACC-EQ|ACC||EQ||-1||1|2|-1|ACC LIMITED|1|5|1|1|1|1|2|0|451510|0|0|0|0|0|225755|0|225755|-451510|2|451510|0|0|0|225755|0|225755^"
// @Failure 400 {string} Token "250|2|0018|Error"
// @Router /CGI/PositionReport [post]
func welcomeController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	body := "<h1>Hello Carrot!</h1>"
	fmt.Fprint(ctx.Writer, body)
}

// PositionReport godoc
// @Summary Position Report
// @Tags Position Report API V1
// @Description Provides position, trade, order and other reports
// @Param request body string false "201|A599059|A599059|26F638C8EADE24F38A6480D7EB53DA||||"
// @Success 200 {string} Token "251|0|0000306|nse_cm|1594|A599059|CNC|INFY-EQ|INFY||EQ||-1||1|2|-1|INFOSYS LIMITED|1|5|1|1|1|1|1|0|171555|0|0|0|0|0|171555|0|171555|-171555|1|171555|0|0|0|171555|0|171555^nse_cm|22|A599059|CNC|ACC-EQ|ACC||EQ||-1||1|2|-1|ACC LIMITED|1|5|1|1|1|1|2|0|451510|0|0|0|0|0|225755|0|225755|-451510|2|451510|0|0|0|225755|0|225755^"
// @Failure 400 {string} Token "250|2|0018|Error"
// @Router /CGI/PositionReport [post]
func latencyController(ctx *gin.Context) {
	log.Debug(ctx).Msg("received report request")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Header().Set("Content-Type", "image/png")
	data, timeSeries := business.HandleMetricsLatency()
	graph := charts.DrawChart(data, timeSeries, "time", "latency")
	graph.Render(chart.PNG, ctx.Writer)
}
