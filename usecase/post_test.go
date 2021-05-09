package usecase

import (
	"context"
	"testing"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	mock_service "github.com/akubi0w1/golang-sample/mock/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPost_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	type in struct {
		accountID entity.AccountID
		title     string
		body      string
		tags      []string
		imageIDs  []int
	}
	tests := []struct {
		name     string
		injector func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage)
		in       in
		out      entity.Post
		code     code.Code
	}{
		{
			name: "failed to get author",
			injector: func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage) {
				mu.EXPECT().GetByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{}, code.Error(code.NotFound, "some error"))
			},
			in: in{
				accountID: "accid",
				title:     "title001",
				body:      "body001",
				tags:      []string{"tag001", "tag002"},
				imageIDs:  []int{1, 2},
			},
			out:  entity.Post{},
			code: code.NotFound,
		},
		{
			name: "failed to create or get tags",
			injector: func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage) {
				mu.EXPECT().GetByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().CreateOrGetMultiple(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{}, code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				title:     "title001",
				body:      "body001",
				tags:      []string{"tag001", "tag002"},
				imageIDs:  []int{1, 2},
			},
			out:  entity.Post{},
			code: code.Database,
		},
		{
			name: "failed to get images",
			injector: func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage) {
				mu.EXPECT().GetByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().CreateOrGetMultiple(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				}, nil)
				mi.EXPECT().GetByIDs(ctx, []int{1, 2}).Return(entity.ImageList{}, code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				title:     "title001",
				body:      "body001",
				tags:      []string{"tag001", "tag002"},
				imageIDs:  []int{1, 2},
			},
			out:  entity.Post{},
			code: code.Database,
		},
		{
			name: "failed to create post",
			injector: func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage) {
				mu.EXPECT().GetByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().CreateOrGetMultiple(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				}, nil)
				mi.EXPECT().GetByIDs(ctx, []int{1, 2}).Return(entity.ImageList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mp.EXPECT().Create(ctx, "title001", "body001", entity.UserID(1), gomock.Any(), gomock.Any()).Return(entity.Post{}, code.Error(code.Database, "some error"))
			},
			in: in{
				accountID: "accid",
				title:     "title001",
				body:      "body001",
				tags:      []string{"tag001", "tag002"},
				imageIDs:  []int{1, 2},
			},
			out:  entity.Post{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mp *mock_service.MockPost, mu *mock_service.MockUser, mt *mock_service.MockTag, mi *mock_service.MockImage) {
				mu.EXPECT().GetByAccountID(ctx, entity.AccountID("accid")).Return(entity.User{
					ID:        1,
					AccountID: "accid",
				}, nil)
				mt.EXPECT().CreateOrGetMultiple(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				}, nil)
				mi.EXPECT().GetByIDs(ctx, []int{1, 2}).Return(entity.ImageList{
					{ID: 1},
					{ID: 2},
				}, nil)
				mp.EXPECT().Create(ctx, "title001", "body001", entity.UserID(1), gomock.Any(), gomock.Any()).Return(entity.Post{
					ID:    1,
					Title: "title001",
					Body:  "body001",
					Tags: entity.TagList{
						{ID: 1, Tag: "tag001"},
						{ID: 2, Tag: "tag002"},
					},
					Images: entity.ImageList{
						{ID: 1},
						{ID: 2},
					},
					AuthorID: 1,
				}, nil)
			},
			in: in{
				accountID: "accid",
				title:     "title001",
				body:      "body001",
				tags:      []string{"tag001", "tag002"},
				imageIDs:  []int{1, 2},
			},
			out: entity.Post{
				ID:    1,
				Title: "title001",
				Body:  "body001",
				Tags: entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				},
				Images: entity.ImageList{
					{ID: 1},
					{ID: 2},
				},
				AuthorID: 1,
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
			mp := mock_service.NewMockPost(ctrl)
			mu := mock_service.NewMockUser(ctrl)
			mt := mock_service.NewMockTag(ctrl)
			mi := mock_service.NewMockImage(ctrl)
			tt.injector(mp, mu, mt, mi)
			uc := NewPost(mp, mu, mt, mi)
			out, err := uc.Create(ctx, tt.in.accountID, tt.in.title, tt.in.body, tt.in.tags, tt.in.imageIDs)
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
		injector func(mp *mock_service.MockPost)
		out      out
		code     code.Code
	}{
		{
			name: "failed to get",
			injector: func(mp *mock_service.MockPost) {
				mp.EXPECT().GetAll(ctx).Return(entity.PostList{}, 0, code.Error(code.Database, "some error"))
			},
			out: out{
				posts: entity.PostList{},
				total: 0,
			},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mp *mock_service.MockPost) {
				mp.EXPECT().GetAll(ctx).Return(entity.PostList{
					{ID: 1},
					{ID: 2},
				}, 2, nil)
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
			mp := mock_service.NewMockPost(ctrl)
			tt.injector(mp)
			uc := NewPost(mp, nil, nil, nil)
			out, total, err := uc.GetAll(ctx)
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
		injector func(mp *mock_service.MockPost)
		in       int
		out      entity.Post
		code     code.Code
	}{
		{
			name: "failed to get",
			injector: func(mp *mock_service.MockPost) {
				mp.EXPECT().GetByID(ctx, 1).Return(entity.Post{}, code.Error(code.NotFound, "some error"))
			},
			in:   1,
			out:  entity.Post{},
			code: code.NotFound,
		},
		{
			name: "success",
			injector: func(mp *mock_service.MockPost) {
				mp.EXPECT().GetByID(ctx, 1).Return(entity.Post{
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
			mp := mock_service.NewMockPost(ctrl)
			tt.injector(mp)
			uc := NewPost(mp, nil, nil, nil)
			out, err := uc.GetByID(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
