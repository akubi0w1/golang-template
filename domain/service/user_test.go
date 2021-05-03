package service

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	mock_repository "github.com/akubi0w1/golang-sample/mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetAll(t *testing.T) {
	ctx := context.Background()
	type out struct {
		users entity.UserList
		total int
	}
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser)
		out      out
		code     code.Code
	}{
		{
			name: "failed to get users",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindAll(ctx).Return(entity.UserList{}, code.Error(code.Database, "some error"))
			},
			out: out{
				users: entity.UserList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "failed to get total",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindAll(ctx).Return(entity.UserList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mu.EXPECT().Count(ctx).Return(0, code.Error(code.Database, "some error"))
			},
			out: out{
				users: entity.UserList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindAll(ctx).Return(entity.UserList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mu.EXPECT().Count(ctx).Return(2, nil)
			},
			out: out{
				users: entity.UserList{
					{ID: 1},
					{ID: 2},
				},
				total: 2,
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu)
			srv := NewUser(mu, mh)
			outUsers, total, err := srv.GetAll(ctx)
			assert.Equal(t, tt.out.users, outUsers)
			assert.Equal(t, tt.out.total, total)
			assert.Equal(t, tt.code, code.GetCode(err))

		})
	}
}

func TestUser_GetByID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser)
		in       entity.UserID
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to get user",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByID(ctx, entity.UserID(1)).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in:   1,
			out:  entity.User{},
			code: code.NotFound,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByID(ctx, entity.UserID(1)).Return(entity.User{
					ID: 1,
				}, nil)
			},
			in: 1,
			out: entity.User{
				ID: 1,
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu)
			srv := NewUser(mu, mh)
			out, err := srv.GetByID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))

		})
	}
}

func TestUser_GetByAccountID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser)
		in       entity.AccountID
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to get user",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in:   "accid",
			out:  entity.User{},
			code: code.NotFound,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
			},
			in: "accid",
			out: entity.User{
				ID:        1,
				AccountID: "accid",
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu)
			srv := NewUser(mu, mh)
			out, err := srv.GetByAccountID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))

		})
	}
}

func TestUser_Create(t *testing.T) {
	ctx := context.Background()
	type in struct {
		accountID entity.AccountID
		email     entity.Email
		password  string
		name      string
		avatarURL string
	}
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser, mh *mock_repository.MockHash)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to generate hash password",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mh.EXPECT().GenerateHashPassword("pass").Return("", code.Error(code.Password, "some error"))
			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.Password,
		},
		{
			name: "failed to check duplicate",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mh.EXPECT().GenerateHashPassword("pass").Return("hash", nil)
				mu.EXPECT().CheckDuplicate(ctx, entity.AccountID("accid"), entity.Email("email")).Return(code.Error(code.BadRequest, "some error"))
			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.BadRequest,
		},
		{
			name: "failed to new user",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mh.EXPECT().GenerateHashPassword("pass").Return("hash", nil)
				mu.EXPECT().CheckDuplicate(ctx, entity.AccountID("acc"), entity.Email("email")).Return(nil)

			},
			in: in{
				accountID: "acc",
				email:     "email",
				password:  "pass",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.BadRequest,
		},
		{
			name: "failed to insert",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mh.EXPECT().GenerateHashPassword("pass").Return("hash", nil)
				mu.EXPECT().CheckDuplicate(ctx, entity.AccountID("accid"), entity.Email("email")).Return(nil)
				mu.EXPECT().Insert(ctx, gomock.Any()).Return(entity.User{}, code.Error(code.Database, "some error"))

			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mh.EXPECT().GenerateHashPassword("pass").Return("hash", nil)
				mu.EXPECT().CheckDuplicate(ctx, entity.AccountID("accid"), entity.Email("email")).Return(nil)
				mu.EXPECT().Insert(ctx, gomock.Any()).Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Email:     "email",
					Password:  "hash",
					Profile: entity.Profile{
						ID:        1,
						Name:      "name",
						AvatarURL: "avatar",
					},
				}, nil)

			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
				name:      "name",
				avatarURL: "avatar",
			},
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Email:     "email",
				Password:  "hash",
				Profile: entity.Profile{
					ID:        1,
					Name:      "name",
					AvatarURL: "avatar",
				},
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu, mh)
			srv := NewUser(mu, mh)
			out, err := srv.Create(ctx, tt.in.accountID, tt.in.email, tt.in.password, tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))

		})
	}
}

func TestUser_UpdateProfile(t *testing.T) {
	ctx := context.Background()
	dummyDate := time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC)
	type in struct {
		accountID entity.AccountID
		name      string
		avatarURL string
	}
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to find user",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in: in{
				accountID: "accid",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.NotFound,
		},
		{
			name: "failed to update",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mu.EXPECT().UpdateProfile(ctx, gomock.Any()).Return(code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				name:      "name",
				avatarURL: "avatar",
			},
			out:  entity.User{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mu.EXPECT().UpdateProfile(ctx, gomock.Any()).Return(nil)
			},
			in: in{
				accountID: "accid",
				name:      "name",
				avatarURL: "avatar",
			},
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				UpdatedAt: dummyDate,
				Profile: entity.Profile{
					Name:      "name",
					AvatarURL: "avatar",
				},
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu)
			srv := NewUser(mu, mh)
			out, err := srv.UpdateProfile(ctx, tt.in.accountID, tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_Delete(t *testing.T) {
	ctx := context.Background()
	dummyDate := time.Date(2021, 5, 3, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser)
		in       entity.AccountID
		code     code.Code
	}{
		{
			name: "failed to find user",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in:   "accid",
			code: code.NotFound,
		},
		{
			name: "failed to set deleted at",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
					DeletedAt: &dummyDate,
				}, nil)
			},
			in:   "accid",
			code: code.BadRequest,
		},
		{
			name: "failed to set delete",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mu.EXPECT().Delete(ctx, gomock.Any()).Return(code.Error(code.Database, "some error"))
			},
			in:   "accid",
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mu.EXPECT().Delete(ctx, gomock.Any()).Return(nil)
			},
			in:   "accid",
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			patch := monkey.Patch(time.Now, func() time.Time { return dummyDate })
			defer patch.Unpatch()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu)
			srv := NewUser(mu, mh)
			err := srv.Delete(ctx, tt.in)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_Authorize(t *testing.T) {
	ctx := context.Background()
	type in struct {
		accountID entity.AccountID
		password  string
	}
	tests := []struct {
		name     string
		injector func(mu *mock_repository.MockUser, mh *mock_repository.MockHash)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to find user",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out:  entity.User{},
			code: code.NotFound,
		},
		{
			name: "failed to validate",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Password:  "hash",
				}, nil)
				mh.EXPECT().ValidatePassword("hash", "pass").Return(code.Error(code.Unauthorized, "some error"))
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out:  entity.User{},
			code: code.Unauthorized,
		},
		{
			name: "success",
			injector: func(mu *mock_repository.MockUser, mh *mock_repository.MockHash) {
				mu.EXPECT().FindByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Password:  "hash",
				}, nil)
				mh.EXPECT().ValidatePassword("hash", "pass").Return(nil)
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Password:  "hash",
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mu := mock_repository.NewMockUser(ctrl)
			mh := mock_repository.NewMockHash(ctrl)
			tt.injector(mu, mh)
			srv := NewUser(mu, mh)
			out, err := srv.Authorize(ctx, tt.in.accountID, tt.in.password)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
