package encryptor

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// BcryptEncryptor is a object that implements a encryptor interface
type BcryptEncryptor struct{}

// CreateHash generate a new hash from string
func (encryptor BcryptEncryptor) CreateHash(aString string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(aString), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Occurred an error while generate a new password hash, err: %s", err)
	}

	return string(hash), nil
}

// Validate a string with a hash
func (encryptor BcryptEncryptor) Validate(aHash, aString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(aHash), []byte(aString))

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
