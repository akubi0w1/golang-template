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

func TestTag_FindByIDs(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       []int
		out      entity.TagList
		code     code.Code
	}{
		{
			name:     "success no ids",
			injector: func(cli *ent.Client) {},
			in:       []int{},
			out:      entity.TagList{},
			code:     code.OK,
		},
		{
			name: "success",
			injector: func(cli *ent.Client) {
				cli.Tag.CreateBulk(
					cli.Tag.Create().SetTag("tag001").SetCreatedAt(dummyTime),
					cli.Tag.Create().SetTag("tag002").SetCreatedAt(dummyTime),
				).Save(ctx)
			},
			in: []int{1, 2},
			out: entity.TagList{
				{ID: 1, Tag: "tag001", CreatedAt: dummyTime},
				{ID: 2, Tag: "tag002", CreatedAt: dummyTime},
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
			u := NewTag(cli)
			out, err := u.FindByIDs(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestTag_FindByTags(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       []string
		out      entity.TagList
		code     code.Code
	}{
		{
			name:     "success no ids",
			injector: func(cli *ent.Client) {},
			in:       []string{},
			out:      entity.TagList{},
			code:     code.OK,
		},
		{
			name: "success",
			injector: func(cli *ent.Client) {
				cli.Tag.CreateBulk(
					cli.Tag.Create().SetTag("tag001").SetCreatedAt(dummyTime),
					cli.Tag.Create().SetTag("tag002").SetCreatedAt(dummyTime),
				).Save(ctx)
			},
			in: []string{"tag001", "tag002"},
			out: entity.TagList{
				{ID: 1, Tag: "tag001", CreatedAt: dummyTime},
				{ID: 2, Tag: "tag002", CreatedAt: dummyTime},
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
			u := NewTag(cli)
			out, err := u.FindByTags(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestTag_InsertBulk(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       entity.TagList
		out      entity.TagList
		code     code.Code
	}{
		{
			name:     "success no ids",
			injector: func(cli *ent.Client) {},
			in:       entity.TagList{},
			out:      entity.TagList{},
			code:     code.OK,
		},
		{
			name:     "success",
			injector: func(cli *ent.Client) {},
			in: entity.TagList{
				{Tag: "tag001", CreatedAt: dummyTime},
				{Tag: "tag002", CreatedAt: dummyTime},
			},
			out: entity.TagList{
				{ID: 1, Tag: "tag001", CreatedAt: dummyTime},
				{ID: 2, Tag: "tag002", CreatedAt: dummyTime},
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
			u := NewTag(cli)
			out, err := u.InsertBulk(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestTag_IsExist(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       string
		out      bool
		code     code.Code
	}{
		{
			name: "success to be not exist",
			injector: func(cli *ent.Client) {
				cli.Tag.CreateBulk(
					cli.Tag.Create().SetTag("tag001").SetCreatedAt(dummyTime),
					cli.Tag.Create().SetTag("tag002").SetCreatedAt(dummyTime),
				).Save(ctx)
			},
			in:   "tag003",
			out:  false,
			code: code.OK,
		},
		{
			name: "success",
			injector: func(cli *ent.Client) {
				cli.Tag.CreateBulk(
					cli.Tag.Create().SetTag("tag001").SetCreatedAt(dummyTime),
					cli.Tag.Create().SetTag("tag002").SetCreatedAt(dummyTime),
				).Save(ctx)
			},
			in:   "tag001",
			out:  true,
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cli := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer cli.Close()
			tt.injector(cli)
			u := NewTag(cli)
			out, err := u.IsExist(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
