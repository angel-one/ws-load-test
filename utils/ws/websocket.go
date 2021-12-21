package ws

import (
	"fmt"
	"github.com/angel-one/ws-load-test/models"
	"github.com/gorilla/websocket"
	"net/url"
)

func CreateSocket(addr string, urlProto string, path string, counter *models.Counter) (*websocket.Conn, error) {
	wsaddr := url.URL{Scheme: urlProto, Host: addr, Path: path}
	c, _, err := websocket.DefaultDialer.Dial(wsaddr.String(), nil)
	counter.Increment()
	if err != nil {
		fmt.Println("Broken WebSocket Conn:", counter.Val)
		counter.HandleFailure()
	} else {
		fmt.Println("Created WebSocket Conn:", counter.Val)
		counter.HandleSuccess()
	}
	return c, err
}
