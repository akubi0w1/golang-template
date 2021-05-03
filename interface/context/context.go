package context

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
)

type key string

const (
	sessionKey   key = "session"
	accountIDKey key = "accountID"
)

func SetToken(ctx context.Context, token entity.Token) context.Context {
	return context.WithValue(ctx, sessionKey, token)
}

func GetToken(ctx context.Context) (entity.Token, error) {
	if ctx.Value(sessionKey) == nil {
		return "", code.Error(code.Context, "failed to get token from context")
	}
	return ctx.Value(sessionKey).(entity.Token), nil
}

func SetAccountID(ctx context.Context, accountID entity.AccountID) context.Context {
	return context.WithValue(ctx, accountIDKey, accountID)
}

func GetAccountID(ctx context.Context) (entity.AccountID, error) {
	if ctx.Value(accountIDKey) == nil {
		return "", code.Error(code.Context, "failed to get accountID from context")
	}
	return ctx.Value(accountIDKey).(entity.AccountID), nil
}

func GetByKeyFromContext(ctx context.Context, key string) (interface{}, error) {
	if ctx.Value(key) == nil {
		return nil, code.Errorf(code.Context, "failed to get value from context: key=%v", key)
	}
	return ctx.Value(key), nil
}
