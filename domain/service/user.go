//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type User interface {
	GetAll(ctx context.Context, opts ...entity.ListOption) (users entity.UserList, total int, err error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	GetByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error)
	Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error)
}

type UserImpl struct {
	user repository.User
}

func NewUser(user repository.User) *UserImpl {
	return &UserImpl{
		user: user,
	}
}

// TODO: add test
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

// TODO: add test
func (srv *UserImpl) GetByID(ctx context.Context, id int) (entity.User, error) {
	return srv.user.FindByID(ctx, id)
}

// TODO: add test
func (srv *UserImpl) GetByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error) {
	return srv.user.FindByAccountID(ctx, accountID)
}

// TODO: add test
func (srv *UserImpl) Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error) {
	// TODO: password hash
	hash := "hash password"

	if err := srv.user.Validate(ctx, accountID, email); err != nil {
		return entity.User{}, err
	}

	user, err := entity.NewUserWithProfile(accountID, email, hash, name, avatarURL)
	if err != nil {
		return entity.User{}, err
	}
	newID, err := srv.user.Insert(ctx, user)
	if err != nil {
		return entity.User{}, err
	}
	user.SetID(newID)

	return user, nil
}
