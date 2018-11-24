package domain

// CahtID uniquely identifies the room
type CahtID uint64

// RoomRepository use case
type RoomRepository interface {
	FindRoomByID(id uint64) []Room
}

// Room what user is in the room
type Room struct {
	ID CahtID
	UserID
}
