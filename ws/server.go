package ws

import (
	"fmt"
	"net/http"

	"github.com/Krapiy/noData-chat-API/usecases"
	"github.com/gorilla/websocket"
)

// StartServer run server with websocket connections
func StartServer(uc usecases.UserDelivery) error {

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
		go target(conn, uc)
	})

	return http.ListenAndServe(":1604", nil)
}
