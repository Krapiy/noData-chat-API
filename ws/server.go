package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// StartServer run server with websocket connections
func StartServer() error {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":%s}`, err), http.StatusInternalServerError)
			return
		}
		go targetController(conn)
	})

	return http.ListenAndServe(":1604", nil)
}
