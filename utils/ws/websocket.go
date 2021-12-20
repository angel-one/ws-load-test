package ws

import (
	"fmt"
	"github.com/angel-one/ws-load-test/models"
	"github.com/gorilla/websocket"
	"net/url"
)

type completion struct {
	body struct {
		code           string
		fileType       string
		line           int
		column         int
		wordToComplete string
		offset         int
	}
}

func CreateSocket(addr string, urlProto string, path string, counter *models.Counter) (*websocket.Conn, error) {
	wsaddr := url.URL{Scheme: urlProto, Host: addr, Path: path}
	c, _, err := websocket.DefaultDialer.Dial(wsaddr.String(), nil)
	counter.Increment()
	if err != nil {
		fmt.Println("Broken WebSocket Conn:", counter.Val)
		counter.Failure()
	} else {
		fmt.Println("Created WebSocket Conn:", counter.Val)
		counter.Success()
	}
	fmt.Println("Success and Failures", counter.Success(), counter.Failure())
	return c, err
}
