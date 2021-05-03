//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

type Hash interface {
	GenerateHashPassword(password string) (string, error)
	ValidatePassword(hashPassword, rawPassword string) error
}
