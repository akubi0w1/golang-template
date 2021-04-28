package config

import "os"

var (
	SessionName string = ""
	Env         string = ""
)

const (
	defaultSessionName = "_default_session"
	defaultEnv         = "development"
)

func init() {
	Env = os.Getenv("ENV")
	if !IsProduct() {
		Env = defaultEnv
	}

	SessionName = os.Getenv("SESSION_NAME")
	if SessionName == "" {
		SessionName = defaultSessionName
	}
}

// TODO: add test
func IsProduct() bool {
	return Env == "product"
}

// TODO: add test
func IsDevelopment() bool {
	return Env == "development"
}
