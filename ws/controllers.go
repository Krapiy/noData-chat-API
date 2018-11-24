package ws

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/usecases"
	"github.com/gorilla/websocket"
)

// controller lists
const (
	userSaltByName = "userSaltByName"
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

func target(conn *websocket.Conn, uc usecases.UserDelivery) {
	for {
		message := &littleRequestRNI{}
		err := conn.ReadJSON(message)
		if err != nil {
			conn.WriteJSON(&littleErrorRNI{err.Error()})
			return
		}

		switch message.Method {
		case userSaltByName:
			hash, e := uc.GetUserEncryptSalt(message.Data[0].(string))
			err = e
			if err == nil {
				resp := littleResponseRNI{hash}
				conn.WriteJSON(resp)
			}
			fmt.Println(err)
		default:
			err = fmt.Errorf("method %s not avaible", message.Method)
		}
		fmt.Println("error", err)
		if err != nil {
			conn.WriteJSON(&littleErrorRNI{err.Error()})
		}
	}
}
