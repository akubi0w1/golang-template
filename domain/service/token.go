package service

import (
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type TokenManager interface {
	Generate(accountID entity.AccountID) (entity.Token, error)
	// ?Validate()
}

type TokenManagerImpl struct {
	token repository.JWT
}

func NewTokenManager(token repository.JWT) *TokenManagerImpl {
	return &TokenManagerImpl{
		token: token,
	}
}

func (tm *TokenManagerImpl) Generate(accountID entity.AccountID) (entity.Token, error) {
	claims, err := entity.NewClaims(accountID)
	if err != nil {
		return "", err
	}
	token, err := tm.token.Generate(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
