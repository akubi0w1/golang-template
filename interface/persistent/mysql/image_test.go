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

func TestImage_FindByIDs(t *testing.T) {
	ctx := context.Background()
	dummyTime := time.Date(2021, 4, 29, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		injector func(cli *ent.Client)
		in       []int
		out      entity.ImageList
		code     code.Code
	}{
		{
			name:     "success no ids",
			injector: func(cli *ent.Client) {},
			in:       []int{},
			out:      entity.ImageList{},
			code:     code.OK,
		},
		{
			name: "success",
			injector: func(cli *ent.Client) {
				u, _ := cli.User.Create().SetAccountID("accid").SetEmail("email").SetPassword("pass").SetCreatedAt(dummyTime).SetUpdatedAt(dummyTime).Save(ctx)
				cli.Image.CreateBulk(
					cli.Image.Create().SetURL("url001").SetCreatedAt(dummyTime).SetUploadedBy(u),
					cli.Image.Create().SetURL("url002").SetCreatedAt(dummyTime).SetUploadedBy(u),
				).Save(ctx)
			},
			in: []int{1, 2},
			out: entity.ImageList{
				{ID: 1, URL: "url001", CreatedBy: 1},
				{ID: 2, URL: "url002", CreatedBy: 1},
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
			u := NewImage(cli)
			out, err := u.FindByIDs(ctx, tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
