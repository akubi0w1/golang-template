package service

import (
	"context"
	"testing"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	mock_repository "github.com/akubi0w1/golang-sample/mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPost_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	type in struct {
		title    string
		body     string
		authorID entity.UserID
		tags     entity.TagList
		images   entity.ImageList
	}
	tests := []struct {
		name     string
		injector func(mp *mock_repository.MockPost)
		in       in
		out      entity.Post
		code     code.Code
	}{
		{
			name:     "failed to create new post",
			injector: func(mp *mock_repository.MockPost) {},
			in: in{
				title:    "",
				body:     "body001",
				authorID: entity.UserID(1),
				tags: entity.TagList{
					{ID: 1},
				},
				images: entity.ImageList{
					{ID: 1},
				},
			},
			out:  entity.Post{},
			code: code.BadRequest,
		},
		{
			name: "failed to insert post",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().Insert(ctx, gomock.Any()).Return(entity.Post{}, code.Error(code.Database, "some error"))
			},
			in: in{
				title:    "title001",
				body:     "body001",
				authorID: entity.UserID(1),
				tags: entity.TagList{
					{ID: 1},
				},
				images: entity.ImageList{
					{ID: 1},
				},
			},
			out:  entity.Post{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().Insert(ctx, gomock.Any()).Return(entity.Post{
					ID:       1,
					Title:    "title001",
					Body:     "body001",
					AuthorID: 1,
					Tags: entity.TagList{
						{ID: 1},
					},
					Images: entity.ImageList{
						{ID: 1},
					},
				}, nil)
			},
			in: in{
				title:    "title001",
				body:     "body001",
				authorID: entity.UserID(1),
				tags: entity.TagList{
					{ID: 1},
				},
				images: entity.ImageList{
					{ID: 1},
				},
			},
			out: entity.Post{
				ID:       1,
				Title:    "title001",
				Body:     "body001",
				AuthorID: 1,
				Tags: entity.TagList{
					{ID: 1},
				},
				Images: entity.ImageList{
					{ID: 1},
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
			mp := mock_repository.NewMockPost(ctrl)
			tt.injector(mp)
			srv := NewPost(mp)
			out, err := srv.Create(ctx, tt.in.title, tt.in.body, tt.in.authorID, tt.in.tags, tt.in.images)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestPost_GetAll(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	type out struct {
		posts entity.PostList
		total int
	}
	tests := []struct {
		name     string
		injector func(mp *mock_repository.MockPost)
		out      out
		code     code.Code
	}{
		{
			name: "failed to get posts",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().FindAll(ctx).Return(entity.PostList{}, code.Error(code.Database, "some error"))
			},
			out: out{
				posts: entity.PostList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "failed to get total",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().FindAll(ctx).Return(entity.PostList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mp.EXPECT().Count(ctx).Return(0, code.Error(code.Database, "some error"))
			},
			out: out{
				posts: entity.PostList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().FindAll(ctx).Return(entity.PostList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mp.EXPECT().Count(ctx).Return(2, nil)
			},
			out: out{
				posts: entity.PostList{
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
			mp := mock_repository.NewMockPost(ctrl)
			tt.injector(mp)
			srv := NewPost(mp)
			out, total, err := srv.GetAll(ctx)
			assert.Equal(t, tt.out.posts, out)
			assert.Equal(t, tt.out.total, total)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestPost_GetByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mp *mock_repository.MockPost)
		in       int
		out      entity.Post
		code     code.Code
	}{
		{
			name: "failed to find post",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().FindByID(ctx, 1).Return(entity.Post{}, code.Error(code.NotFound, "some error"))
			},
			in:   1,
			out:  entity.Post{},
			code: code.NotFound,
		},
		{
			name: "success",
			injector: func(mp *mock_repository.MockPost) {
				mp.EXPECT().FindByID(ctx, 1).Return(entity.Post{
					ID: 1,
				}, nil)

			},
			in: 1,
			out: entity.Post{
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
			mp := mock_repository.NewMockPost(ctrl)
			tt.injector(mp)
			srv := NewPost(mp)
			out, err := srv.GetByID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
