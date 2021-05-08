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

func TestTag_CreateOrGetMultiple(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mt *mock_repository.MockTag)
		in       []string
		out      entity.TagList
		code     code.Code
	}{
		{
			name: "failed to check being exist",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag").Return(false, code.Error(code.Database, "some error"))
			},
			in:   []string{"tag"},
			out:  entity.TagList{},
			code: code.Database,
		},
		{
			name: "failed to new tag",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag001").Return(false, nil)
				mt.EXPECT().IsExist(ctx, "").Return(false, nil)
			},
			in:   []string{"tag001", ""},
			out:  entity.TagList{},
			code: code.BadRequest,
		},
		{
			name: "failed to insert",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag001").Return(false, nil)
				mt.EXPECT().InsertBulk(ctx, gomock.Any()).Return(entity.TagList{}, code.Error(code.Database, "some error"))
			},
			in:   []string{"tag001"},
			out:  entity.TagList{},
			code: code.Database,
		},
		{
			name: "failed to find all",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag001").Return(false, nil)
				mt.EXPECT().IsExist(ctx, "tag002").Return(true, nil)
				mt.EXPECT().InsertBulk(ctx, gomock.Any()).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
				}, nil)
				mt.EXPECT().FindByTags(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{}, code.Error(code.Database, "some error"))
			},
			in:   []string{"tag001", "tag002"},
			out:  entity.TagList{},
			code: code.Database,
		},
		{
			name: "success to no create",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag001").Return(true, nil)
				mt.EXPECT().IsExist(ctx, "tag002").Return(true, nil)
				mt.EXPECT().FindByTags(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				}, nil)
			},
			in: []string{"tag001", "tag002"},
			out: entity.TagList{
				{ID: 1, Tag: "tag001"},
				{ID: 2, Tag: "tag002"},
			},
			code: code.OK,
		},
		{
			name: "success to create or get",
			injector: func(mt *mock_repository.MockTag) {
				mt.EXPECT().IsExist(ctx, "tag001").Return(false, nil)
				mt.EXPECT().IsExist(ctx, "tag002").Return(true, nil)
				mt.EXPECT().InsertBulk(ctx, gomock.Any()).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
				}, nil)
				mt.EXPECT().FindByTags(ctx, []string{"tag001", "tag002"}).Return(entity.TagList{
					{ID: 1, Tag: "tag001"},
					{ID: 2, Tag: "tag002"},
				}, nil)
			},
			in: []string{"tag001", "tag002"},
			out: entity.TagList{
				{ID: 1, Tag: "tag001"},
				{ID: 2, Tag: "tag002"},
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
			mt := mock_repository.NewMockTag(ctrl)
			tt.injector(mt)
			srv := NewTag(mt)
			out, err := srv.CreateOrGetMultiple(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
