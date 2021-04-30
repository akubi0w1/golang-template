//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type User interface {
	GetAll(ctx context.Context, opts ...entity.ListOption) (users entity.UserList, total int, err error)
}

type UserImpl struct {
	user repository.User
}

func NewUser(user repository.User) *UserImpl {
	return &UserImpl{
		user: user,
	}
}

// TODO-akubi: add test
func (srv *UserImpl) GetAll(ctx context.Context, opts ...entity.ListOption) (entity.UserList, int, error) {
	users, err := srv.user.FindAll(ctx, opts...)
	if err != nil {
		return entity.UserList{}, 0, err
	}
	total, err := srv.user.Count(ctx)
	if err != nil {
		return entity.UserList{}, 0, err
	}
	return users, total, nil
}
