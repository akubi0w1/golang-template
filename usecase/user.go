//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package usecase

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/service"
)

type User interface {
	GetAll(ctx context.Context, opts ...entity.ListOption) (users entity.UserList, total int, err error)
	GetByID(ctx context.Context, id entity.UserID) (entity.User, error)
	Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password string) (entity.User, error)
	CreateWithProfile(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error)
}

type UserImpl struct {
	user service.User
}

func NewUser(user service.User) *UserImpl {
	return &UserImpl{
		user: user,
	}
}

// TODO: test
func (us *UserImpl) GetAll(ctx context.Context, opts ...entity.ListOption) (entity.UserList, int, error) {
	users, total, err := us.user.GetAll(ctx, opts...)
	if err != nil {
		return entity.UserList{}, 0, err
	}
	return users, total, nil
}

// TODO: test
func (us *UserImpl) GetByID(ctx context.Context, id entity.UserID) (entity.User, error) {
	return us.user.GetByID(ctx, id)
}

// TODO: test
func (us *UserImpl) Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password string) (entity.User, error) {
	return us.user.Create(ctx, accountID, email, password, "", "")
}

// TODO: test
func (us *UserImpl) CreateWithProfile(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error) {
	return us.user.Create(ctx, accountID, email, password, name, avatarURL)
}
