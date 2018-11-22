package usecases

import (
	"fmt"

	"github.com/Krapiy/noData-chat-API/domain"
	"github.com/pkg/errors"
)

// UserDelivery display methods for access for UI
type UserDelivery interface {
	GetUserEncryptSalt(string) (string, error)
}

// UserInteractor uses cases for user
type UserInteractor struct {
	UserRepository domain.UserRepository
}

// GetUserEncryptSalt get user salt encrypt user pubkey
func (i *UserInteractor) GetUserEncryptSalt(name string) (string, error) {
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
