package business

import (
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/ws"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

func receiveMsg(wsconn *websocket.Conn, done chan *models.Routine, rout *models.Routine) {
	for {
		_, message, err := wsconn.ReadMessage()
		rout.ReceiveTime = time.Now()
		rout.Diff = rout.ReceiveTime.Sub(rout.SendTime)
		rout.ReceivedMsg = string(message)
		if err != nil {
			log.Println("read:", err)
			return
		}
		done <- rout
	}
}

func writeMsg(wsconn *websocket.Conn, base *models.Base, rout *models.Routine) {
	time.Sleep(time.Second * time.Duration(base.Delay))
	rout.SendTime = time.Now()
	wsconn.WriteMessage(websocket.TextMessage, base.Msg)
}

func singleTest(counter *models.Counter, queue chan *models.Routine, base *models.Base, rout *models.Routine) {
	doneCh := make(chan *models.Routine)
	conn, err := ws.CreateSocket(base.URL, base.Proto, base.Path, counter)
	if err != nil {
		return
	}
	go writeMsg(conn, base, rout)
	go receiveMsg(conn, doneCh, rout)
	queue <- <-doneCh
}

func LoadTest(base *models.Base, latencyCh chan []float64, timeCh chan []time.Time) {

	queue := make(chan *models.Routine, 1)
	globalCounter := &models.Counter{0, sync.Mutex{}, 0, 0}
	localCounter := 0

	var latency []float64
	var timeSeries []time.Time

	for range time.Tick(time.Millisecond * time.Duration(base.TickDelay)) {
		routine := &models.Routine{time.Now(), time.Now(), 0, ""}
		go singleTest(globalCounter, queue, base, routine)
		localCounter++
		if localCounter == base.Count {
			break
		}
	}

	go func() {
		bufferLimit := 0
		for req := range queue {
			latency = append(latency, req.Diff.Seconds()*1000)
			timeSeries = append(timeSeries, req.SendTime)
			bufferLimit++
			if bufferLimit == base.Count {
				latencyCh <- latency
				timeCh <- timeSeries
			}
		}
	}()

}
