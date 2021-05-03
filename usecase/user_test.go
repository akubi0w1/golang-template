package usecase

import (
	"context"
	"testing"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_service "github.com/akubi0w1/golang-sample/mock/service"
)

func TestUser_GetAll(t *testing.T) {
	ctx := context.Background()
	type out struct {
		users entity.UserList
		total int
	}
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser)
		out      out
		code     code.Code
	}{
		{
			name: "failed to get users",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().GetAll(ctx).Return(entity.UserList{}, 0, code.Error(code.Database, "some error"))
			},
			out: out{
				users: entity.UserList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().GetAll(ctx).Return(entity.UserList{
					{ID: 1},
					{ID: 2},
				}, 2, nil)
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			out, total, err := uc.GetAll(ctx)
			assert.Equal(t, tt.out.users, out)
			assert.Equal(t, tt.out.total, total)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_GetByID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser)
		in       entity.UserID
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to get user",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().GetByID(ctx, entity.UserID(1)).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in:   1,
			out:  entity.User{},
			code: code.NotFound,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().GetByID(ctx, entity.UserID(1)).Return(entity.User{
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			out, err := uc.GetByID(ctx, tt.in)
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
	}
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to create",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Create(ctx, entity.AccountID("accid"), entity.Email("email"), "pass", "", "").Return(entity.User{}, code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
			},
			out:  entity.User{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Create(ctx, entity.AccountID("accid"), entity.Email("email"), "pass", "", "").Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Email:     "email",
					Password:  "pass",
					Profile: entity.Profile{
						ID: 1,
					},
				}, nil)
			},
			in: in{
				accountID: "accid",
				email:     "email",
				password:  "pass",
			},
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Email:     "email",
				Password:  "pass",
				Profile: entity.Profile{
					ID: 1,
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			out, err := uc.Create(ctx, tt.in.accountID, tt.in.email, tt.in.password)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_CreateWithProfile(t *testing.T) {
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
		injector func(mu *mock_service.MockUser)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to create",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Create(ctx, entity.AccountID("accid"), entity.Email("email"), "pass", "name", "avatar").Return(entity.User{}, code.Error(code.Database, "some error"))
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
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Create(ctx, entity.AccountID("accid"), entity.Email("email"), "pass", "name", "avatar").Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Email:     "email",
					Password:  "pass",
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
				Password:  "pass",
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			out, err := uc.CreateWithProfile(ctx, tt.in.accountID, tt.in.email, tt.in.password, tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_UpdateProfile(t *testing.T) {
	ctx := context.Background()
	type in struct {
		accountID entity.AccountID
		name      string
		avatarURL string
	}
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser)
		in       in
		out      entity.User
		code     code.Code
	}{
		{
			name: "failed to update",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().UpdateProfile(ctx, entity.AccountID("accid"), "rename", "reavatar").Return(entity.User{}, code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				name:      "rename",
				avatarURL: "reavatar",
			},
			out:  entity.User{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().UpdateProfile(ctx, entity.AccountID("accid"), "rename", "reavatar").Return(entity.User{
					ID:        1,
					AccountID: "accid",
					Email:     "email",
					Password:  "pass",
					Profile: entity.Profile{
						ID:        1,
						Name:      "rename",
						AvatarURL: "reavatar",
					},
				}, nil)
			},
			in: in{
				accountID: "accid",
				name:      "rename",
				avatarURL: "reavatar",
			},
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Email:     "email",
				Password:  "pass",
				Profile: entity.Profile{
					ID:        1,
					Name:      "rename",
					AvatarURL: "reavatar",
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			out, err := uc.UpdateProfile(ctx, tt.in.accountID, tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_Delete(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser)
		in       entity.AccountID
		code     code.Code
	}{
		{
			name: "failed to delete",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Delete(ctx, entity.AccountID("accid")).Return(code.Error(code.Database, "some error"))
			},
			in:   "accid",
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser) {
				mu.EXPECT().Delete(ctx, entity.AccountID("accid")).Return(nil)
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu)
			uc := NewUser(mu, mt)
			err := uc.Delete(ctx, tt.in)
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
	type out struct {
		user  entity.User
		token entity.Token
	}
	tests := []struct {
		name     string
		injector func(mu *mock_service.MockUser, mt *mock_service.MockTokenManager)
		in       in
		out      out
		code     code.Code
	}{
		{
			name: "failed to auth user",
			injector: func(mu *mock_service.MockUser, mt *mock_service.MockTokenManager) {
				mu.EXPECT().Authorize(ctx, entity.AccountID("accid"), "pass").Return(entity.User{}, code.Error(code.Unauthorized, "some error"))
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out: out{
				user:  entity.User{},
				token: "",
			},
			code: code.Unauthorized,
		},
		{
			name: "failed to generate token",
			injector: func(mu *mock_service.MockUser, mt *mock_service.MockTokenManager) {
				mu.EXPECT().Authorize(ctx, entity.AccountID("accid"), "pass").Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().Generate(entity.AccountID("accid")).Return(entity.Token(""), code.Error(code.JWT, "some error"))
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out: out{
				user:  entity.User{},
				token: "",
			},
			code: code.JWT,
		},
		{
			name: "success",
			injector: func(mu *mock_service.MockUser, mt *mock_service.MockTokenManager) {
				mu.EXPECT().Authorize(ctx, entity.AccountID("accid"), "pass").Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().Generate(entity.AccountID("accid")).Return(entity.Token("token"), nil)
			},
			in: in{
				accountID: "accid",
				password:  "pass",
			},
			out: out{
				user: entity.User{
					ID:        1,
					AccountID: "accid",
				},
				token: "token",
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
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTokenManager(ctrl)
			tt.injector(mu, mt)
			uc := NewUser(mu, mt)
			user, token, err := uc.Authorize(ctx, tt.in.accountID, tt.in.password)
			assert.Equal(t, tt.out.user, user)
			assert.Equal(t, tt.out.token, token)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
