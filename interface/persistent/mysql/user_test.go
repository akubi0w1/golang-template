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
