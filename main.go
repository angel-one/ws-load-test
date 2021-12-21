package main

import (
	"fmt"
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/business"
	"github.com/angel-one/ws-load-test/constants"
	_ "github.com/angel-one/ws-load-test/docs"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/configs"
	"github.com/angel-one/ws-load-test/utils/flags"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"runtime"
)

func main() {
	initConfigs()
	startLogger()
	runtime.GOMAXPROCS(runtime.NumCPU())
	results := make(chan *models.TestResult)
	business.LoadTest(results)
	business.HandleMetrices(results)
	StartHTTPServer()
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

func StartHTTPServer() {
	err := http.ListenAndServe(":"+"8090", nil)
	if err != nil {
	}
}
