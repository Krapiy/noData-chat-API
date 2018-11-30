package usecases

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/domain"
	"github.com/pkg/errors"
)

// UserDelivery display methods for access for UI
type UserDelivery interface {
	GetUserEncryptSalt(string) (map[string]interface{}, error)
	GetMessagesByChatID(int) ([]*domain.Message, error)
	InsertMessageByChatID(*domain.Message) (*domain.Message, error)
}

// UserInteractor uses cases for user
type UserInteractor struct {
	UserRepository    domain.UserRepository
	MessageRepository domain.MessageRepository
}

// GetUserEncryptSalt get user salt encrypt user pubkey
func (i *UserInteractor) GetUserEncryptSalt(name string) (map[string]interface{}, error) {
	user, err := i.UserRepository.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("user %s not found", name)
	}

	encryptSalt, err := user.EncryptSalt()
	if err != nil {
		return nil, errors.Wrap(err, "invalid generate salt:")
	}

	keys, err := i.UserRepository.SelectUserPubKeyExcept(name)
	if err != nil {
		return nil, fmt.Errorf("cannot get pub_key")
	}

	res := map[string]interface{}{
		"salt":     encryptSalt,
		"pub_keys": keys,
	}

	return res, nil
}

func (i *UserInteractor) GetMessagesByChatID(id int) ([]*domain.Message, error) {
	messages, err := i.MessageRepository.SelectMessagesByChatID(id)
	if err != nil {
		return nil, fmt.Errorf("cannot get messages by 'room_id': %v", id)
	}
	return messages, nil
}

func (i *UserInteractor) InsertMessageByChatID(message *domain.Message) (*domain.Message, error) {
	message, err := i.MessageRepository.InsertMessageByChatID(message)
	if err != nil {
		return nil, fmt.Errorf("cannot send message")
	}
	return message, nil
}
