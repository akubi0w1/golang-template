//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import (
	"github.com/akubi0w1/golang-sample/domain/entity"
)

type JWT interface {
	Generate(claims entity.Claims) (entity.Token, error)
	Parse(token entity.Token) (entity.Claims, error)
}
