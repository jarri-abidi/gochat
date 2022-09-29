package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/websocket"
)

type Message struct {
	MsgID     int64     `json:"messageId"`
	MsgFrom   int64     `json: "messageFrom"`
	MsgTo     int64     `json: "messageTo"`
	Content   string    `json: "content"`
	CreatedAt time.Time `json: "createdAt"`
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	http.ListenAndServe(":8001", ws(logger))
}

func ws(logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var upgrader websocket.Upgrader
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log("err", err)
			return
		}
		defer conn.Close()

		var msg Message
		for {
			err = conn.ReadJSON(&msg)
			if err != nil {
				logger.Log("err", err)
				return
			}
			logger.Log(
				"from", "client",
				"msgId", msg.MsgID,
				"from", msg.MsgFrom,
				"To", msg.MsgTo,
				"content", msg.Content,
				"createdAt", msg.CreatedAt,
			)

			msg.MsgTo = msg.MsgFrom
			msg.MsgFrom = -1
			if err := conn.WriteJSON(&msg); err != nil {
				logger.Log("err", err)
				return
			}
		}
	}
}
