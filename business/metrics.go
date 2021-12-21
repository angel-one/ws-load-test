package business

import (
	"github.com/angel-one/ws-load-test/models"
	"time"
)

var mainCh chan *models.TestResult

func HandleMetricsLatency() ([]float64, []time.Time) {
	return nil, nil
}

func SetMainChannel (ch chan *models.TestResult){
	mainCh = ch
}
