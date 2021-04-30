package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/enttest"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestUser_FindAll(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		out      entity.UserList
		code     code.Code
	}{
		{
			name:     "success no user",
			injector: func(cli *ent.Client) {},
			out:      entity.UserList{},
			code:     code.OK,
		},
		{
			name: "success has profile",
			injector: func(cli *ent.Client) {
				p, _ := cli.Profile.Create().SetName("name").SetAvatarURL("avatar").Save(ctx)
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetProfile(p).Save(ctx)
			},
			out: entity.UserList{
				{
					ID:        1,
					AccountID: "accid",
					Email:     "email",
					Password:  "pass",
					CreatedAt: dummyTime,
					UpdatedAt: dummyTime,
					Profile: entity.Profile{
						ID:        1,
						Name:      "name",
						AvatarURL: "avatar",
					},
				},
			},
			code: code.OK,
		},
		{
			name: "success deleted is nil",
			injector: func(cli *ent.Client) {
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).Save(ctx)
			},
			out:  entity.UserList{},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewUser(cli)
			out, err := u.FindAll(ctx)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_FindByAccountID(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       entity.AccountID
		out      entity.User
		code     code.Code
	}{
		{
			name:     "no user",
			injector: func(cli *ent.Client) {},
			out:      entity.User{},
			in:       "accid",
			code:     code.NotFound,
		},
		{
			name: "success has profile",
			injector: func(cli *ent.Client) {
				p, _ := cli.Profile.Create().SetName("name").SetAvatarURL("avatar").Save(ctx)
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetProfile(p).Save(ctx)
			},
			in: "accid",
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Email:     "email",
				Password:  "pass",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				Profile: entity.Profile{
					ID:        1,
					Name:      "name",
					AvatarURL: "avatar",
				},
			},
			code: code.OK,
		},
		{
			name: "success deleted is nil",
			injector: func(cli *ent.Client) {
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).Save(ctx)
			},
			in:   "accid",
			out:  entity.User{},
			code: code.NotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewUser(cli)
			out, err := u.FindByAccountID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_FindByID(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       entity.UserID
		out      entity.User
		code     code.Code
	}{
		{
			name:     "no user",
			injector: func(cli *ent.Client) {},
			out:      entity.User{},
			in:       1,
			code:     code.NotFound,
		},
		{
			name: "success has profile",
			injector: func(cli *ent.Client) {
				p, _ := cli.Profile.Create().SetName("name").SetAvatarURL("avatar").Save(ctx)
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetProfile(p).Save(ctx)
			},
			in: 1,
			out: entity.User{
				ID:        1,
				AccountID: "accid",
				Email:     "email",
				Password:  "pass",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				Profile: entity.Profile{
					ID:        1,
					Name:      "name",
					AvatarURL: "avatar",
				},
			},
			code: code.OK,
		},
		{
			name: "success deleted is nil",
			injector: func(cli *ent.Client) {
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).Save(ctx)
			},
			in:   1,
			out:  entity.User{},
			code: code.NotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewUser(cli)
			out, err := u.FindByID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_Insert(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       entity.User
		out      entity.User
		code     code.Code
	}{
		{
			name:     "success",
			injector: func(cli *ent.Client) {},
			in: entity.User{
				AccountID: "accid",
				Password:  "pass",
				Email:     "emal",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				Profile: entity.Profile{
					Name:      "name",
					AvatarURL: "ava",
				},
			},
			out: entity.User{
				ID:        entity.UserID(1),
				AccountID: "accid",
				Password:  "pass",
				Email:     "emal",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				Profile: entity.Profile{
					ID:        1,
					Name:      "name",
					AvatarURL: "ava",
				},
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewUser(cli)
			out, err := u.Insert(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_Count(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		out      int
		code     code.Code
	}{
		{
			name:     "no user",
			injector: func(cli *ent.Client) {},
			out:      0,
			code:     code.OK,
		},
		{
			name: "success",
			injector: func(cli *ent.Client) {
				p, _ := cli.Profile.Create().SetName("name").SetAvatarURL("avatar").Save(ctx)
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetProfile(p).Save(ctx)
			},
			out:  1,
			code: code.OK,
		},
		{
			name: "success deleted is nil",
			injector: func(cli *ent.Client) {
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).Save(ctx)
			},
			out:  0,
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewUser(cli)
			out, err := u.Count(ctx)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
