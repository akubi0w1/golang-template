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
)

func TestPost_Insert(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       entity.Post
		out      entity.Post
		code     code.Code
	}{
		{
			name: "success to post",
			injector: func(cli *ent.Client) {
				cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
			},
			in: entity.Post{
				Title:     "title001",
				Body:      "body001",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				AuthorID:  entity.UserID(1),
			},
			out: entity.Post{
				ID:        1,
				Title:     "title001",
				Body:      "body001",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				AuthorID:  entity.UserID(1),
				Tags:      entity.TagList{},
				Images:    entity.ImageList{},
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
			srv := NewPost(cli)
			out, err := srv.Insert(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestPost_FindAll(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		out      entity.PostList
		code     code.Code
	}{
		{
			name: "success",
			injector: func(cli *ent.Client) {
				u, _ := cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
				cli.Post.CreateBulk(
					cli.Post.Create().SetTitle("title001").SetBody("body001").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetAuthor(u),
					cli.Post.Create().SetTitle("title002").SetBody("body002").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).SetAuthor(u),
				).Save(ctx)
			},
			out: entity.PostList{
				{
					ID:        1,
					Title:     "title001",
					Body:      "body001",
					CreatedAt: dummyTime,
					UpdatedAt: dummyTime,
					AuthorID:  entity.UserID(1),
					Tags:      entity.TagList{},
					Images:    entity.ImageList{},
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
			srv := NewPost(cli)
			out, err := srv.FindAll(ctx)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestPost_FindByID(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       int
		out      entity.Post
		code     code.Code
	}{
		{
			name: "success",
			injector: func(cli *ent.Client) {
				u, _ := cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
				cli.Post.Create().SetTitle("title001").SetBody("body001").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetAuthor(u).Save(ctx)
			},
			in: 1,
			out: entity.Post{
				ID:        1,
				Title:     "title001",
				Body:      "body001",
				CreatedAt: dummyTime,
				UpdatedAt: dummyTime,
				AuthorID:  entity.UserID(1),
				Tags:      entity.TagList{},
				Images:    entity.ImageList{},
			},
			code: code.OK,
		},
		{
			name: "success deletedAt is not null",
			injector: func(cli *ent.Client) {
				u, _ := cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
				cli.Post.Create().SetTitle("title001").SetBody("body001").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).SetAuthor(u).Save(ctx)
			},
			in:   1,
			out:  entity.Post{},
			code: code.NotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			srv := NewPost(cli)
			out, err := srv.FindByID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestPost_Count(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		out      int
		code     code.Code
	}{
		{
			name: "success",
			injector: func(cli *ent.Client) {
				u, _ := cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
				cli.Post.CreateBulk(
					cli.Post.Create().SetTitle("title001").SetBody("body001").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetAuthor(u),
					cli.Post.Create().SetTitle("title002").SetBody("body002").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).SetDeletedAt(dummyTime).SetAuthor(u),
				).Save(ctx)
			},
			out:  1,
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			srv := NewPost(cli)
			out, err := srv.Count(ctx)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
