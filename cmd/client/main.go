package main

import (
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/websocket"
)

type Message struct {
	MsgID     int64  `json: "message_id"`
	MsgFrom   int64  `json: "message_from"`
	MsgTo     int64  `json: "message_to"`
	Content   string `json: "content"`
	CreatedAt string `json: "created_at"`
}

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
		var msg Message
		for {
			conn.ReadJSON(&msg)
			if err != nil {
				logger.Log("err", err)
				return
			}
			logger.Log("from", "server", "msgId", msg.MsgId, "from", msg.MsgFrom, "To", msg.MsgTo, "content", msg.Content, "createdAt", msg.CreatedAt)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	var idcounter int64
	var msg Message
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			idcounter += 1
			// will be changed later
			msg.MsgId = idcounter
			msg.MsgFrom = 1
			msg.MsgTo = -1
			msg.Content = "hi"
			msg.CreatedAt = t.String()
			err := conn.WriteJSON(&msg)
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
