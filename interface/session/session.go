package session

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/config"
	"github.com/akubi0w1/golang-sample/interface/jwt"
)

// Session is session
type Session struct {
	accountID string
}

func NewSession(accountID string) *Session {
	return &Session{
		accountID: accountID,
	}
}

func (s *Session) SetAccountID(accountID string) {
	s.accountID = accountID
}

func (s *Session) GetAccountID() string {
	return s.accountID
}

func (s *Session) IsEmpty() bool {
	return s.accountID == ""
}

// SessionManage is manage session
type SessionManager interface {
	Start(w http.ResponseWriter, session Session) (string, error)
	Restart(w http.ResponseWriter, r *http.Request, session Session) (string, error)
	End(w http.ResponseWriter, r *http.Request) error
	Get(r *http.Request) (session Session, err error)
}

type SessionManagerImpl struct {
	tokenHandler jwt.TokenHandler
}

func NewSessionManager() *SessionManagerImpl {
	return &SessionManagerImpl{
		tokenHandler: jwt.NewJWTHandler(),
	}
}

func (s *SessionManagerImpl) Start(w http.ResponseWriter, ss Session) (string, error) {
	tokenString, err := s.tokenHandler.Generate(ss.accountID)
	if err != nil {
		return "", err
	}
	setCookie(w, tokenString)
	return tokenString, nil
}

func (s *SessionManagerImpl) Restart(w http.ResponseWriter, r *http.Request, ss Session) (string, error) {
	if err := deleteCookie(w, r, config.SessionName); err != nil {
		return "", err
	}
	return s.Start(w, ss)
}

func (s *SessionManagerImpl) End(w http.ResponseWriter, r *http.Request) error {
	return deleteCookie(w, r, config.SessionName)
}

func (s *SessionManagerImpl) Get(r *http.Request) (ss Session, err error) {
	cookie, err := r.Cookie(config.SessionName)
	if err != nil {
		return Session{}, code.Errorf(code.Session, "failed to get cookie: %v", err)
	}
	claims, err := s.tokenHandler.Parse(cookie.Value)
	if err != nil {
		return Session{}, err
	}
	return Session{
		accountID: claims.GetAccountID(),
	}, nil
}

func setCookie(w http.ResponseWriter, value string) http.ResponseWriter {
	cookie := &http.Cookie{
		Name:     config.SessionName,
		Value:    value,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	return w
}

func getCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, code.Errorf(code.Session, "failed to get cookie: %s", err)
	}
	return cookie, nil
}

func deleteCookie(w http.ResponseWriter, r *http.Request, cookieName string) error {
	cookie, err := getCookie(r, cookieName)
	if err != nil {
		return code.Errorf(code.Session, "failed to delete cookie: %s", err)
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	return nil
}
