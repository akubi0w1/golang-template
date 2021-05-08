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

func TestImage_GetByIDs(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tests := []struct {
		name     string
		injector func(mi *mock_repository.MockImage)
		in       []int
		out      entity.ImageList
		code     code.Code
	}{
		{
			name: "failed to find images",
			injector: func(mi *mock_repository.MockImage) {
				mi.EXPECT().FindByIDs(ctx, []int{1, 2}).Return(entity.ImageList{}, code.Error(code.Database, "some error"))
			},
			in:   []int{1, 2},
			out:  entity.ImageList{},
			code: code.Database,
		},
		{
			name: "success",
			injector: func(mi *mock_repository.MockImage) {
				mi.EXPECT().FindByIDs(ctx, []int{1, 2}).Return(entity.ImageList{
					{ID: 1},
					{ID: 2},
				}, nil)
			},
			in: []int{1, 2},
			out: entity.ImageList{
				{ID: 1},
				{ID: 2},
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
			mi := mock_repository.NewMockImage(ctrl)
			tt.injector(mi)
			srv := NewImage(mi)
			out, err := srv.GetByIDs(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
