package strategy

import (
	"github.com/angel-one/go-utils/log"
	"github.com/angel-one/ws-load-test/models"
	"github.com/angel-one/ws-load-test/utils/flags"
	"github.com/gorilla/websocket"
	"time"
)

func HandleBasic(wsconn *websocket.Conn, testState *models.TestResult) {
	done := make(chan struct{})
	sendTime := time.Now()
	isSendDone := false
	go func() {
		defer close(done)
		for {
			_, message, err := wsconn.ReadMessage()
			testState.ReceiveTimeLatest = time.Now()
			testState.ReceivedMsgCount++
			if isSendDone {
				testState.Latency = testState.EventTime.Second() - sendTime.Second()
				isSendDone = false
			}
			if err != nil {
				log.Error(nil).Err(err).Msg("error receiving message from ws")
				return
			}
			log.Info(nil).Str("message", string(message)).Msg("received message from ws")
		}
	}()
	ticker := time.NewTicker(time.Second * time.Duration(flags.LifeTime()))
	defer ticker.Stop()
	for {
		time.Sleep(time.Second * time.Duration(flags.WriteTime()))
		testState.SendTimeLatest = time.Now()
		testState.SendMsgCount++
		isSendDone = true
		sendTime = testState.EventTime
		wsconn.WriteMessage(websocket.TextMessage, []byte(flags.MessageText()))
		select {
		case <-done:
			log.Info(nil).Msg("closing ws")
			return
		case <-ticker.C:
			log.Info(nil).Msg("closing ws post life time")
			return
		}
	}

}
