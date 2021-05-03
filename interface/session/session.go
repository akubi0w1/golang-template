package session

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/config"
	"github.com/akubi0w1/golang-sample/domain/entity"
)

func Start(w http.ResponseWriter, token entity.Token) {
	setCookie(w, token.String())
}

func setCookie(w http.ResponseWriter, value string) http.ResponseWriter {
	cookie := &http.Cookie{
		Name:     config.SessionCookieName,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	return w
}
