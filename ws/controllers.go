package ws

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/domain"

	"github.com/Krapiy/noData-chat-API/usecases"
	"github.com/gorilla/websocket"
)

// controller lists
const (
	getEncryptInfo      = "getEncryptInfo"
	getMessagesByRoomID = "getMessagesByRoomID"
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
		case getEncryptInfo:
			info, e := uc.GetEncryptInfo(message.Data[0].(string))
			err = e
			if err == nil {
				resp := littleResponseRNI{info}
				conn.WriteJSON(resp)
			}
		case getMessagesByRoomID:
			messages, e := uc.GetMessagesByRoomID(int(message.Data[0].(float64)))
			err = e
			if err == nil {
				resp := littleResponseRNI{messages}
				conn.WriteJSON(resp)
			}
		case sendMessageToRoom:
			userID := domain.UserID(message.Data[1].(float64))
			message := &domain.Message{
				RoomID:       int(message.Data[0].(float64)),
				UserSenderID: userID,
				Message:      message.Data[2].(string),
			}
			newMessage, e := uc.InsertMessageByRoomID(message)
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
