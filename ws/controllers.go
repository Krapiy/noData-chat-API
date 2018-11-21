package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// controller lists
const (
	getUserByName = "getUserByName"
)

type littleRequestRNI struct {
	Method string        `json:"method"`
	Data   []interface{} `json:"data"`
}

type littleResponseRNI struct {
	Result interface{} `json:"result"`
}

type littleErrorRNI struct {
	Error string `json:"error"`
}

func targetController(conn *websocket.Conn) {
	for {
		message := &littleRequestRNI{}
		err := conn.ReadJSON(message)
		if err != nil {
			conn.WriteJSON(&littleErrorRNI{err.Error()})
			return
		}

		switch message.Method {
		case getUserByName:
			err = conn.WriteJSON(&littleResponseRNI{"pub_key"})
		default:
			err = fmt.Errorf("method %s not avaible", message.Method)
		}

		if err != nil {
			conn.WriteJSON(&littleErrorRNI{err.Error()})
		}
	}
}
