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

// StartSession starts session with cookie
// func StartSession(handler http.Handler) http.Handler {
// 	logger := log.New()
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		handler.ServeHTTP(w, r)

// 		// TODO: なんとかして、ここにtokenを持ってくる必要がある
// 		ctx := r.Context()

// 		token, err := context.GetToken(ctx)
// 		if err != nil {
// 			logger.Error("failed to start session: %v", err)
// 			return
// 		}
// 		logger.Debug("aaaaaaaaaaa", token)

// 		setCookie(w, token.String())

// 	}
// 	return http.HandlerFunc(fn)
// }

// func setCookie(w http.ResponseWriter, value string) http.ResponseWriter {
// 	cookie := &http.Cookie{
// 		Name:     config.SessionCookieName,
// 		Value:    value,
// 		HttpOnly: true,
// 		Secure:   true,
// 		Path:     "/",
// 	}
// 	http.SetCookie(w, cookie)
// 	return w
// }

// func getCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
// 	cookie, err := r.Cookie(cookieName)
// 	if err != nil {
// 		return nil, code.Errorf(code.Session, "failed to get cookie: %s", err)
// 	}
// 	return cookie, nil
// }

// func deleteCookie(w http.ResponseWriter, r *http.Request, cookieName string) error {
// 	cookie, err := getCookie(r, cookieName)
// 	if err != nil {
// 		return code.Errorf(code.Session, "failed to delete cookie: %s", err)
// 	}
// 	cookie.MaxAge = -1
// 	http.SetCookie(w, cookie)
// 	return nil
// }
