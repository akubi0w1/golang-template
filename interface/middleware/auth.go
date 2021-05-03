package middleware

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/config"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/context"
	"github.com/akubi0w1/golang-sample/interface/jwt"
	"github.com/akubi0w1/golang-sample/interface/response"
)

// SaveSessionToContext get and save session to context
func SaveSessionToContext(handler http.Handler) http.Handler {
	tokenManager := jwt.NewJWTImpl()
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookie, err := r.Cookie(config.SessionCookieName)
		if err != nil {
			handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		claims, err := tokenManager.Parse(entity.Token(cookie.Value))
		if err != nil {
			handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ctx = context.SetAccountID(ctx, claims.AccountID)

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// Authorize check session is not empty
func Authorize(handler http.Handler) http.Handler {
	tokenManager := jwt.NewJWTImpl()
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie(config.SessionCookieName)
		if err != nil {
			response.Error(w, r, code.Errorf(code.Unauthorized, "failed to get cookie: %v", err))
			return
		}

		claims, err := tokenManager.Parse(entity.Token(cookie.Value))
		if err != nil {
			response.Error(w, r, err)
			return
		}

		ctx = context.SetAccountID(ctx, claims.AccountID)

		handler.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
