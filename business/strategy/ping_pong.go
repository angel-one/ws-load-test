package strategy

import (
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/flags"
	"github.com/gorilla/websocket"
	"time"
)

func HandlePingPong(wsconn *websocket.Conn, testState *models.TestResult, queue chan *models.TestResult) {
	done := make(chan struct{})
	sendTime := time.Now()
	isSendDone := false
	go func() {
		defer close(done)
		for {
			_, message, err := wsconn.ReadMessage()
			testState.ReceiveTimeLatest = time.Now()
			testState.ReceivedMsgCount++
			if testState.ReceivedMsgCount == 1{
				testState.ReceiveTimeFirst = time.Now()
			}
			if isSendDone {
				testState.Latency = testState.EventTime.Second() - sendTime.Second()
				isSendDone = false
			}
			testState.EventType = "receive"
			testState.EventTime = time.Now()
			if err != nil {
				log.Error(nil).Err(err).Msg("error receiving message from ws")
				return
			}
			sendResult := testState
			queue <- sendResult
			log.Info(nil).Str("message", string(message)).Msg("received message from ws")
		}
	}()
	ticker := time.NewTicker(time.Second * time.Duration(flags.LifeTime()))
	defer ticker.Stop()
	for {
		time.Sleep(time.Second * time.Duration(flags.WriteTime()))
		testState.SendTimeLatest = time.Now()
		testState.SendMsgCount++
		if testState.SendMsgCount == 1{
			testState.SendTimeFirst =  time.Now()
		}
		testState.EventType = "send"
		isSendDone = true
		sendTime = testState.EventTime
		testState.EventTime = time.Now()
		wsconn.WriteMessage(websocket.TextMessage, []byte("ping"))
		log.Info(nil).Str("message", string("ping")).Msg("sent message to ws")
		sendResult := testState
		queue <- sendResult
		select {
		case <-done:
			log.Info(nil).Msg("closing ws")
			testState.HasEnded = true
			testState.EndTime = time.Now()
			testState.EventType = "done"
			testState.EventTime = time.Now()
			sendResult := testState
			queue <- sendResult
			return
		case <-ticker.C:
			log.Info(nil).Msg("closing ws post life time")
			testState.HasEnded = true
			testState.EndTime = time.Now()
			testState.EventType = "done"
			testState.EventTime = time.Now()
			sendResult := testState
			queue <- sendResult
			return
		}
	}
}
