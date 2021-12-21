package main

import (
	"fmt"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/business"
	"github.com/angel-one/ws-load-test/constants"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/chart"
	"github.com/angel-one/ws-load-test/utils/configs"
	"github.com/angel-one/ws-load-test/utils/flags"
	"net/http"
	"runtime"
	"time"

	_ "github.com/angel-one/ws-load-test/docs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initConfigs()
	startLogger()
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)
	currentTest := &models.Base{flags.Host(), flags.Protocol(), flags.Request(),
		[]byte("ABC"), flags.WriteTime(), flags.HoldTime(), flags.Path()}
	business.LoadTest(currentTest, latency, timeSeries)

	data := <-latency
	timeData := <-timeSeries
	fmt.Println(data, timeData)
	fmt.Println("Running HTTP Server, Check /latency route at Port 8900")
	StartHTTPServer("8900", data, timeData)
	fmt.Scanln()
}

func initConfigs() {
	configs.Init(flags.BaseConfigPath())
}

func startLogger() {
	loggerConfig, err := configs.Get(constants.LoggerConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting logger config")
	}
	log.InitLogger(log.Level(loggerConfig.GetString(constants.LogLevelConfigKey)))
}

func StartHTTPServer(port string, latency []float64, timeSeries []time.Time) {
	http.HandleFunc("/", chart.RenderHTML)
	http.HandleFunc("/latency", func(w http.ResponseWriter, r *http.Request) {
		chart.DrawChart(w, r, latency, timeSeries)
	})
	fmt.Printf("HTTP Server Listening at... %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
	}
}
