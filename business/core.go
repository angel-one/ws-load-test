package business

import (
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/business/strategy"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/flags"
	"github.com/angel-one/ws-load-test/utils/ws"
	"sync"
	"time"
)

func test(counter *models.Counter, queue chan *models.TestResult, result *models.TestResult) {
	conn, err := ws.CreateSocket(flags.Host(), flags.Protocol(), flags.Path(), counter)
	defer conn.Close()
	if err != nil {
		log.Error(nil).Err(err).Msg("not able to create a ws connection")
		result.HasEnded = true
		result.HasError = true
		result.EndTime = time.Now()
		result.EventType = "error"
		result.EventTime = time.Now()
		sendResult  := result
		queue <- sendResult
		return
	}
	if flags.Strategy() != "" {
		if flags.Strategy() == "ping_pong" {
			strategy.HandlePingPong(conn, result, queue)
		} else if flags.Strategy() == "exchange_tick" {
			strategy.HandleExchangeTick(conn, result)
		}
	} else {
		strategy.HandleBasic(conn, result, queue)
	}
}

func LoadTest(queue chan *models.TestResult) {

	globalCounter := &models.Counter{0, sync.Mutex{}, 0, 0}
	counter := 0
	for range time.Tick(time.Millisecond * time.Duration(flags.GapTime())) {
		counter++
		routine := &models.TestResult{}
		routine.ID = int64(counter)
		routine.StartTime = time.Now()
		routine.HasEnded = false
		routine.HasError = false
		routine.Latency = 0
		routine.EventTime = time.Now()
		routine.EventType = "start"
		go test(globalCounter, queue, routine)
		if counter == flags.Request() {
			break
		}
	}
}
