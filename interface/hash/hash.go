package hash

import (
	"github.com/akubi0w1/golang-sample/code"
	"golang.org/x/crypto/bcrypt"
)

type HashImpl struct {
}

func NewHash() *HashImpl {
	return &HashImpl{}
}

func (h *HashImpl) GenerateHashPassword(password string) (string, error) {
	if len(password) > 71 {
		return "", code.Error(code.Password, "password length too large")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", code.Errorf(code.Password, "failed to generata hash: %v", err)
	}
	return string(hash), nil
}

func (h *HashImpl) ValidatePassword(hashPassword, rawPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword)); err != nil {
		return code.Errorf(code.Unauthorized, "failed to validate password: %v", err)
	}
	return nil
}
