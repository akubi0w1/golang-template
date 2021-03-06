//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type User interface {
	GetAll(ctx context.Context, opts ...entity.ListOption) (users entity.UserList, total int, err error)
	GetByID(ctx context.Context, id entity.UserID) (entity.User, error)
	GetByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error)
	Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error)
	UpdateProfile(ctx context.Context, accountID entity.AccountID, name string, avatarURL string) (entity.User, error)
	Delete(ctx context.Context, accountID entity.AccountID) error
	Authorize(ctx context.Context, accountID entity.AccountID, password string) (entity.User, error)
}

type UserImpl struct {
	user repository.User
	hash repository.Hash
}

func NewUser(user repository.User, hash repository.Hash) *UserImpl {
	return &UserImpl{
		user: user,
		hash: hash,
	}
}

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

func (srv *UserImpl) GetByID(ctx context.Context, id entity.UserID) (entity.User, error) {
	return srv.user.FindByID(ctx, id)
}

func (srv *UserImpl) GetByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error) {
	return srv.user.FindByAccountID(ctx, accountID)
}

func (srv *UserImpl) Create(ctx context.Context, accountID entity.AccountID, email entity.Email, password, name, avatarURL string) (entity.User, error) {
	hash, err := srv.hash.GenerateHashPassword(password)
	if err != nil {
		return entity.User{}, err
	}

	if err := srv.user.CheckDuplicate(ctx, accountID, email); err != nil {
		return entity.User{}, err
	}

	user, err := entity.NewUserWithProfile(accountID, email, hash, name, avatarURL)
	if err != nil {
		return entity.User{}, err
	}
	user, err = srv.user.Insert(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (srv *UserImpl) UpdateProfile(ctx context.Context, accountID entity.AccountID, name string, avatarURL string) (entity.User, error) {
	user, err := srv.user.FindByAccountID(ctx, accountID)
	if err != nil {
		return entity.User{}, err
	}
	user.UpdateProfile(name, avatarURL)
	if err = srv.user.UpdateProfile(ctx, user); err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (srv *UserImpl) Delete(ctx context.Context, accountID entity.AccountID) error {
	user, err := srv.user.FindByAccountID(ctx, accountID)
	if err != nil {
		return err
	}
	if err = user.Delete(); err != nil {
		return err
	}

	if err = srv.user.Delete(ctx, user); err != nil {
		return err
	}
	return nil
}

func (srv *UserImpl) Authorize(ctx context.Context, accountID entity.AccountID, password string) (entity.User, error) {
	user, err := srv.user.FindByAccountID(ctx, accountID)
	if err != nil {
		return entity.User{}, err
	}
	if err = srv.hash.ValidatePassword(user.Password, password); err != nil {
		return entity.User{}, err
	}
	return user, nil
}
