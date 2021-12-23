package main

import (
	"fmt"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/go-utils/middlewares"
	"github.com/angel-one/ws-load-test/api"
	"github.com/angel-one/ws-load-test/business"
	"github.com/angel-one/ws-load-test/constants"
	_ "github.com/angel-one/ws-load-test/docs"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/configs"
	"github.com/angel-one/ws-load-test/utils/flags"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

// @title WS Load Test
// @version 1.0
// @description Load test of websockets
// @termsOfService https://swagger.io/terms/
// @contact.name Team AMX
// @contact.email AmxTechTeamInternal@angelbroking.com
// @BasePath /
func main() {
	initConfigs()
	startLogger()
	runtime.GOMAXPROCS(runtime.NumCPU())
	results := make(chan *models.TestResult)
	business.SetMainChannel(results)
	go business.Init()
	go startRouter()
	business.LoadTest(results)
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

func startRouter() {
	router := api.GetRouter(middlewares.Logger(middlewares.LoggerMiddlewareOptions{}))
	err := router.Run(fmt.Sprintf(":%d", flags.ServerPort()))
	if err != nil {
		log.Fatal(nil).Err(err).Msg("error starting router")
	}
	log.Info(nil).Msg("router started")
}
