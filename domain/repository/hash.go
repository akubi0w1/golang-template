package repository

type Hash interface {
	GenerateHashPassword(password string) (string, error)
	ValidatePassword(hashPassword, rawPassword string) error
}
