package domain

// MessageRepository  use case
type MessageRepository interface {
	SelectMessagesByChatID(int) ([]*Message, error)
	InsertMessageByChatID(*Message) (*Message, error)
}

// Message display user send message
type Message struct {
	ID             int    `json:"id" db:"id"`
	UserSenderID   UserID `json:"user_sender_id" db:"user_sender_id"`
	UserReceiverID UserID `json:"user_receiver_id" db:"user_receiver_id"`
	Message        string `json:"message" db:"message"`
	CahtID         int    `json:"room_id" db:"room_id"`
}
