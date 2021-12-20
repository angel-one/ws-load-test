package main

import (
	"fmt"
	"github.com/angel-one/ws-load-test/api"
	"github.com/angel-one/ws-load-test/constants"
	"github.com/angel-one/ws-load-test/utils/configs"
	"github.com/angel-one/ws-load-test/utils/flags"
	"github.com/angel-one/ws-load-test/utils/httpclient"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"runtime"
	"time"

	_ "github.com/angel-one/ws-load-test/docs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initConfigs()
	startLogger()
	initHTTPClient()
	runtime.GOMAXPROCS(runtime.NumCPU())
	latency := make(chan []float64)
	timeSeries := make(chan []time.Time)
//test
	data := <-latency
	timeData := <-timeSeries

	startRouter()
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

func initHTTPClient() {
	applicationConfig, err := configs.Get(constants.ApplicationConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting application config")
	}

	err = httpclient.Init(httpclient.Config{
		ConnectTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPConnectTimeoutInMillisKey),
		KeepAliveDuration: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPKeepAliveDurationInMillisKey),
		MaxIdleConnections: applicationConfig.GetInt(constants.HTTPMaxIdleConnectionsKey),
		IdleConnectionTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPIdleConnectionTimeoutInMillisKey),
		TLSHandshakeTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPTlsHandshakeTimeoutInMillisKey),
		ExpectContinueTimeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPExpectContinueTimeoutInMillisKey),
		Timeout: time.Millisecond *
			applicationConfig.GetDuration(constants.HTTPTimeoutInMillisKey),
	})
	if err != nil {
		log.Fatal(nil).Err(err).Msg("unable to initialize http client")
	}
}

func startRouter() {
	// get router
	router := api.GetRouter(middlewares.Logger(middlewares.LoggerMiddlewareOptions{}))
	// now start router
	err := router.Run(fmt.Sprintf(":%d", flags.Port()))
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error starting router")
	}
}
