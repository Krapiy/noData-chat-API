package domain

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const soldSize = 64

// Config use for user settings
// type Config struct {}

// UserID uniquely identifies the user
type UserID uint64

// UserRepository use case
type UserRepository interface {
	FindByName(string) (*User, error)
}

// User constatis all info about user
type User struct {
	ID       int    `db:"id"`
	Name     string `db:"user_name"`
	Password string `db:"password_hash"`
	PubKey   string `db:"pub_key"`
	// Config
}

// EncryptSalt encrypt random salt
func (u *User) EncryptSalt() (string, error) {
	salt := make([]byte, soldSize)
	if _, err := rand.Read(salt); err != nil {
		return "", errors.Wrap(err, "generate sold error:")
	}

	raw := strings.Split(u.PubKey, " ")
	key := make([]byte, len(raw))
	for r, v := range raw {
		k, _ := strconv.Atoi(v)
		key[r] = byte(k)
	}

	encryptSalt, err := rsaEncrypt(salt, key)
	if err != nil {
		return "", errors.Wrapf(err, "encrypt salt:")
	}

	return string(encryptSalt), nil
}

func rsaEncrypt(data, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid prase pubkey")
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}
