package main

import (
	"fmt"
	"github.com/angel-one/go-example-project/api"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-example-project/utils/configs"
	"github.com/angel-one/go-example-project/utils/flags"
	"github.com/angel-one/go-example-project/utils/httpclient"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"time"

	_ "github.com/angel-one/go-example-project/docs"
	_ "github.com/go-sql-driver/mysql"
)

// @title Go Example Project
// @version 1.0
// @description Go Example Project
// @termsOfService https://swagger.io/terms/

// @contact.name Shubham Sinha
// @contact.email shubham.sinha@angelbroking.com

// @BasePath /

func main() {
	initConfigs()
	startLogger()
	initHTTPClient()
	startRouter()
}

func initConfigs() {
	// init configs
	configs.Init(flags.BaseConfigPath())
}

func startLogger() {
	// start logger
	loggerConfig, err := configs.Get(constants.LoggerConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting logger config")
	}
	log.InitLogger(log.Level(loggerConfig.GetString(constants.LogLevelConfigKey)))
}

func initHTTPClient() {
	// get application configs
	applicationConfig, err := configs.Get(constants.ApplicationConfig)
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error getting application config")
	}

	// init http client
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
