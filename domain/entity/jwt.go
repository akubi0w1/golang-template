package entity

import "github.com/akubi0w1/golang-sample/code"

type Token string

func (t Token) String() string {
	return string(t)
}

type Claims struct {
	AccountID
}

func NewClaims(accountID AccountID) (Claims, error) {
	if accountID == "" {
		return Claims{}, code.Error(code.Session, "accountID as claims is empty")
	}
	return Claims{
		AccountID: accountID,
	}, nil
}
