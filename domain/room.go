package domain

// RoomID uniquely identifies the room
type RoomID uint64

// RoomRepository use case
type RoomRepository interface {
	FindRoomByID(id uint64) []Room
}

// Room what user is in the room
type Room struct {
	ID RoomID
	UserID
}
