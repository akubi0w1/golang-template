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

// TODO: add test
func NewUser(accountID AccountID, email Email, password string) (User, error) {
	if err := accountID.validateAccountID(); err != nil {
		return User{}, err
	}
	now := time.Now()
	return User{
		AccountID: accountID,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
		Profile:   Profile{},
	}, nil
}

// TODO: add test
func NewUserWithProfile(accountID AccountID, email Email, password, name, avatarURL string) (User, error) {
	if err := accountID.validateAccountID(); err != nil {
		return User{}, err
	}
	now := time.Now()
	return User{
		AccountID: accountID,
		Email:     email,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
		Profile: Profile{
			Name:      name,
			AvatarURL: avatarURL,
		},
	}, nil
}

// TODO: add test
func (u *User) SetID(id int) {
	u.ID = id
}

type UserList []User

type AccountID string

// TODO: add test
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

// TODO: add test
func (e Email) String() string {
	return string(e)
}

type Profile struct {
	ID        int
	Name      string
	AvatarURL string
}
