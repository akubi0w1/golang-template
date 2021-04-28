package entity

type User struct {
	ID        int
	AccountID AccountID
	Email     Email
	Password  string
}

func NewUser(accountID AccountID, email Email, password string) User {
	return User{
		AccountID: accountID,
		Email:     email,
		Password:  password,
	}
}

type UserList []User

type AccountID string

type Email string
