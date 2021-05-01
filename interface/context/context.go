package context

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/interface/session"
)

type key string

const (
	sessionKey key = "session"
)

func SetSession(ctx context.Context, ss session.Session) context.Context {
	return context.WithValue(ctx, sessionKey, ss)
}

func GetSessionFromContext(ctx context.Context) (session.Session, error) {
	if ctx.Value(sessionKey) == nil {
		return session.Session{}, code.Errorf(code.Context, "failed to get session from context")
	}
	return ctx.Value(sessionKey).(session.Session), nil
}

func GetByKeyFromContext(ctx context.Context, key string) (interface{}, error) {
	if ctx.Value(key) == nil {
		return nil, code.Errorf(code.Context, "failed to get value from context: key=%v", key)
	}
	return ctx.Value(key), nil
}
