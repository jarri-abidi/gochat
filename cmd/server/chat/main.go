package main

import (
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/gorilla/websocket"
)

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

		for {
			mt, msg, err := conn.ReadMessage()
			if err != nil {
				logger.Log("err", err)
				return
			}
			logger.Log("type", mt, "from", "client", "msg", string(msg))

			if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
				logger.Log("err", err)
				return
			}
		}
	}
}
