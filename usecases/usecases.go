package usecases

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/domain"
	"github.com/pkg/errors"
)

// UserDelivery display methods for access for UI
type UserDelivery interface {
	GetEncryptInfo(string) (map[string]interface{}, error)
	GetMessagesByRoomID(int) ([]*domain.Message, error)
	InsertMessageByRoomID(*domain.Message) (*domain.Message, error)
}

// UserInteractor uses cases for user
type UserInteractor struct {
	UserRepository    domain.UserRepository
	MessageRepository domain.MessageRepository
}

func (i *UserInteractor) getUserEncryptSalt(name string) (string, error) {
	user, err := i.UserRepository.FindByName(name)
	if err != nil {
		return "", fmt.Errorf("user %s not found", name)
	}

	encryptSalt, err := user.EncryptSalt()
	if err != nil {
		return "", errors.Wrap(err, "invalid generate salt:")
	}

	return encryptSalt, nil
}

func (i *UserInteractor) getPubKeysExceptTarget(name string) ([]*domain.User, error) {
	keys, err := i.UserRepository.SelectUsersPubKeyExcept(name)
	if err != nil {
		return nil, fmt.Errorf("cannot get pub_key")
	}
	return keys, nil
}

// GetEncryptInfo get user salt encrypt user pubkey
func (i *UserInteractor) GetEncryptInfo(name string) (map[string]interface{}, error) {

	encryptSalt, err := i.getUserEncryptSalt(name)
	if err != nil {
		return nil, err
	}

	keys, err := i.getPubKeysExceptTarget(name)
	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"salt":     encryptSalt,
		"pub_keys": keys,
	}

	return res, nil
}

func (i *UserInteractor) GetMessagesByRoomID(id int) ([]*domain.Message, error) {
	messages, err := i.MessageRepository.SelectMessagesByRoomID(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get messages by 'room_id': %v", id)
	}
	return messages, nil
}

func (i *UserInteractor) InsertMessageByRoomID(message *domain.Message) (*domain.Message, error) {
	message, err := i.MessageRepository.InsertMessageByRoomID(message)
	if err != nil {
		return nil, fmt.Errorf("cannot send message")
	}
	return message, nil
}
