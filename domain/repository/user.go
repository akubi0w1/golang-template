//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
)

type User interface {
	FindAll(ctx context.Context, opts ...entity.ListOption) (entity.UserList, error)
	FindByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error)
	FindByID(ctx context.Context, id entity.UserID) (entity.User, error)
	Insert(ctx context.Context, user entity.User) (entity.User, error)
	Count(ctx context.Context) (int, error)
	Validate(ctx context.Context, accountID entity.AccountID, email entity.Email) error
}
