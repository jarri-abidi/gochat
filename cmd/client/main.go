package main

import (
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/websocket"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	u := url.URL{Scheme: "ws", Host: "localhost:8001"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logger.Log("err", err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				logger.Log("err", err)
				return
			}
			logger.Log("type", mt, "from", "server", "msg", string(msg))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte("hi"))
			if err != nil {
				logger.Log("err", err)
				return
			}
		case <-interrupt:
			logger.Log("msg", "interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logger.Log("err", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
