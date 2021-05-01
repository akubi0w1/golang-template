package mysql

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	entuser "github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/user"
)

type UserImpl struct {
	cli *ent.Client
}

func NewUser(cli *ent.Client) *UserImpl {
	return &UserImpl{
		cli: cli,
	}
}

func (u *UserImpl) FindAll(ctx context.Context, opts ...entity.ListOption) (entity.UserList, error) {
	opt := mergeListOptions(opts)

	users, err := u.cli.User.Query().
		Where(entuser.DeletedAtIsNil()).
		Order(ent.Asc(entuser.FieldID)).
		Limit(opt.Limit).
		Offset(opt.Offset).
		WithProfile().
		All(ctx)
	if err != nil {
		return entity.UserList{}, code.Errorf(code.Database, "failed to find all user: %v", err)
	}
	return toEntityUsers(users), nil
}

func (u *UserImpl) FindByAccountID(ctx context.Context, accountID entity.AccountID) (entity.User, error) {
	user, err := u.cli.User.Query().
		Where(
			entuser.And(
				entuser.AccountID(accountID.String()),
				entuser.DeletedAtIsNil(),
			),
		).
		WithProfile().
		Only(ctx)
	if err != nil {
		return entity.User{}, code.Errorf(code.NotFound, "failed to find user by accountID=%s: %v", accountID, err)
	}
	return toEntityUser(user), nil
}

func (u *UserImpl) FindByID(ctx context.Context, id entity.UserID) (entity.User, error) {
	user, err := u.cli.User.Query().
		Where(
			entuser.And(
				entuser.IDEQ(id.Int()),
				entuser.DeletedAtIsNil(),
			),
		).
		WithProfile().
		Only(ctx)
	if err != nil {
		return entity.User{}, code.Errorf(code.NotFound, "failed to find user by ID=%d: %v", id, err)
	}
	return toEntityUser(user), nil
}

func (u *UserImpl) Insert(ctx context.Context, user entity.User) (entity.User, error) {
	newUser, err := u.cli.User.Create().
		SetAccountID(user.AccountID.String()).
		SetEmail(user.Email.String()).
		SetPassword(user.Password).
		SetCreatedAt(user.CreatedAt).
		SetUpdatedAt(user.UpdatedAt).
		Save(ctx)
	if err != nil {
		return entity.User{}, code.Errorf(code.Database, "failed to insert user: %v", err)
	}
	profile, err := u.cli.Profile.Create().
		SetName(user.Profile.Name).
		SetAvatarURL(user.Profile.AvatarURL).
		SetUser(newUser).
		Save(ctx)
	if err != nil {
		return entity.User{}, code.Errorf(code.Database, "failed to insert profile: %v", err)
	}
	// TODO: これ自動でセットされて欲しいところ...
	newUser.Edges.Profile = profile
	return toEntityUser(newUser), nil
}

func (u *UserImpl) Count(ctx context.Context) (int, error) {
	total, err := u.cli.User.Query().
		Where(entuser.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return 0, code.Errorf(code.Database, "failed to count active user: %v", err)
	}
	return total, nil
}

// TODO: add test
func (u *UserImpl) CheckDuplicate(ctx context.Context, accountID entity.AccountID, email entity.Email) error {
	existAccountID, err := u.cli.User.Query().
		Where(entuser.AccountIDEQ(accountID.String())).
		Exist(ctx)
	if err != nil {
		return code.Errorf(code.Database, "failed to check exist accountID=%s: %v", accountID, err)
	}
	if existAccountID {
		return code.Errorf(code.Conflict, "accountID=%s is conflict", accountID)
	}

	existEamil, err := u.cli.User.Query().
		Where(entuser.EmailEQ(email.String())).
		Exist(ctx)
	if err != nil {
		return code.Errorf(code.Database, "failed to check exist email=%s: %v", email, err)
	}
	if existEamil {
		return code.Errorf(code.Conflict, "email=%s is conflict", email)
	}
	return nil
}

// TODO: add test
func (u *UserImpl) UpdateProfile(ctx context.Context, user entity.User) error {
	_, err := u.cli.Profile.UpdateOneID(user.Profile.ID).
		SetName(user.Profile.Name).
		SetAvatarURL(user.Profile.AvatarURL).
		Save(ctx)
	if err != nil {
		return code.Errorf(code.Database, "failed to update profile by userID=%d: %v", user.ID, err)
	}
	_, err = u.cli.User.UpdateOneID(user.ID.Int()).
		SetUpdatedAt(user.UpdatedAt).
		SetEmail(user.Email.String()).
		Save(ctx)
	if err != nil {
		return code.Errorf(code.Database, "failed to update userID=%d: %v", user.ID, err)
	}
	return nil
}

// TODO: add test
func (u *UserImpl) Delete(ctx context.Context, user entity.User) error {
	_, err := u.cli.User.UpdateOneID(user.ID.Int()).
		SetDeletedAt(*user.DeletedAt).
		Save(ctx)
	if err != nil {
		return code.Errorf(code.Database, "failed to up deleteFlag by userID=%d: %v", user.ID, err)
	}
	return nil
}

func toEntityUser(user *ent.User) entity.User {
	return entity.User{
		ID:        entity.UserID(user.ID),
		AccountID: entity.AccountID(user.AccountID),
		Email:     entity.Email(user.Email),
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Profile:   toEntityProfile(user.Edges.Profile),
	}
}

func toEntityUsers(users []*ent.User) entity.UserList {
	us := make(entity.UserList, 0, len(users))
	for i := range users {
		us = append(us, toEntityUser(users[i]))
	}
	return us
}

func toEntityProfile(profile *ent.Profile) entity.Profile {
	if profile == nil {
		return entity.Profile{}
	}
	return entity.Profile{
		ID:        profile.ID,
		Name:      profile.Name,
		AvatarURL: profile.AvatarURL,
	}
}
