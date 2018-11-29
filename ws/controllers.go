package ws

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/domain"

	"github.com/Krapiy/noData-chat-API/usecases"
	"github.com/gorilla/websocket"
)

// controller lists
const (
	userSaltByName      = "userSaltByName"
	getMessagesByChatID = "getMessagesByChatID"
	sendMessageToRoom   = "sendMessageToRoom"
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
		case getMessagesByChatID:
			messages, e := uc.GetMessagesByChatID(int(message.Data[0].(float64)))
			err = e
			if err == nil {
				resp := littleResponseRNI{messages}
				conn.WriteJSON(resp)
			}
		case sendMessageToRoom:
			userID := domain.UserID(message.Data[1].(float64))
			message := &domain.Message{
				CahtID:       int(message.Data[0].(float64)),
				UserSenderID: userID,
				Message:      message.Data[2].(string),
			}
			newMessage, e := uc.InsertMessageByChatID(message)
			err = e
			if err == nil {
				resp := littleResponseRNI{newMessage}
				conn.WriteJSON(resp)
			}
		default:
			err = fmt.Errorf("method %s not avaible", message.Method)
		}
		if err != nil {
			conn.WriteJSON(&littleErrorRNI{err.Error()})
		}
	}
}
