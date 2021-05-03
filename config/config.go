package config

import "os"

var (
	SessionCookieName string = ""
	Env               string = ""
)

const (
	defaultSessionCookieName = "_session"
	defaultEnv               = "development"
)

func init() {
	Env = os.Getenv("ENV")
	if !IsProduction() {
		Env = defaultEnv
	}

	SessionCookieName = os.Getenv("SESSION_COOKIE_NAME")
	if SessionCookieName == "" {
		SessionCookieName = defaultSessionCookieName
	}
}

// TODO: add test
func IsProduction() bool {
	return Env == "production"
}

// TODO: add test
func IsDevelopment() bool {
	return Env == "development"
}
