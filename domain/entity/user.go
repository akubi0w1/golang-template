package entity

import (
	"time"

	"github.com/akubi0w1/golang-sample/code"
)

type User struct {
	ID        int
	AccountID AccountID
	Email     Email
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Profile   Profile
}

func NewUser(accountID AccountID, email Email, password string) (User, error) {
	if err := accountID.validateAccountID(); err != nil {
		return User{}, err
	}
	return User{
		AccountID: accountID,
		Email:     email,
		Password:  password,
	}, nil
}

type UserList []User

type AccountID string

func (ai AccountID) String() string {
	return string(ai)
}

func (ai AccountID) validateAccountID() error {
	if len(ai) > 20 {
		return code.Error(code.BadRequest, "accountID is too long")
	}
	if len(ai) < 5 {
		return code.Error(code.BadRequest, "accountID is too short")
	}
	return nil
}

type Email string

type Profile struct {
	ID        int
	Name      string
	AvatarURL string
}
