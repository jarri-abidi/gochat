package main

import (
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/websocket"
)

type Message struct {
	MsgID     int64  `json:"message_id"`
	MsgFrom   int64  `json: "message_from"`
	MsgTo     int64  `json: "message_to"`
	Content   string `json: "content"`
	CreatedAt string `json: "created_at"`
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
			conn.ReadJSON(&msg)
			if err != nil {
				logger.Log("err", err)
				return
			}
			logger.Log("msgId", msg.MsgId, "from", msg.MsgFrom, "To", msg.MsgTo, "content", msg.Content, "createdat", msg.CreatedAt)

			msg.MsgTo = msg.MsgFrom
			msg.MsgFrom = -1
			if err := conn.WriteJSON(&msg); err != nil {
				logger.Log("err", err)
				return
			}
		}
	}
}
