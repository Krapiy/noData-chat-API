package domain

// MessageRepository  use case
type MessageRepository interface {
	FindMessageByID(id uint64) Message
	FindMessagesByChanID(id CahtID) []Message
}

// Message display user send message
type Message struct {
	ID             uint64
	UserSenderID   UserID
	UserReceiverID UserID
	CahtID
}
