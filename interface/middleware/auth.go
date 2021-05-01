package middleware

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/interface/context"
	"github.com/akubi0w1/golang-sample/interface/response"
	"github.com/akubi0w1/golang-sample/interface/session"
)

// SaveSessionToContext get and save session to context
func SaveSessionToContext(handler http.Handler) http.Handler {
	sm := session.NewSessionManager()
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ss, _ := sm.Get(r)
		ctx = context.SetSession(ctx, ss)

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorize check session is not empty
func Authorize(handler http.Handler) http.Handler {
	sm := session.NewSessionManager()
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ss, err := sm.Get(r)
		if err != nil {
			response.Error(w, r, err)
			return
		}
		if ss.IsEmpty() {
			response.Error(w, r, code.Error(code.Unauthorized, "session is empty"))
			return
		}
		ctx = context.SetSession(ctx, ss)

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
