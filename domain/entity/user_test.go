package entity

import (
	"strings"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/akubi0w1/golang-sample/code"
	"github.com/stretchr/testify/assert"
)

func TestUser_NewUserWithProfile(t *testing.T) {
	dummyDate := time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC)
	type in struct {
		accountID AccountID
		email     Email
		password  string
		name      string
		avatarURL string
	}
	tests := []struct {
		name string
		in   in
		out  User
		code code.Code
	}{
		{
			name: "failed to accountID too large",
			in: in{
				accountID: AccountID(strings.Repeat("a", 21)),
			},
			out:  User{},
			code: code.BadRequest,
		},
		{
			name: "failed to accountID too short",
			in: in{
				accountID: "a",
			},
			out:  User{},
			code: code.BadRequest,
		},
		{
			name: "success",
			in: in{
				accountID: "accid",
				email:     "hoge@example.com",
				password:  "password",
				name:      "name",
				avatarURL: "avatar",
			},
			out: User{
				AccountID: "accid",
				Email:     "hoge@example.com",
				Password:  "password",
				CreatedAt: dummyDate,
				UpdatedAt: dummyDate,
				Profile: Profile{
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
			patch := monkey.Patch(time.Now, func() time.Time { return dummyDate })
			defer patch.Unpatch()
			out, err := NewUserWithProfile(tt.in.accountID, tt.in.email, tt.in.password, tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUser_UpdateProfile(t *testing.T) {
	dummyDate := time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC)
	dummyUser := User{}
	type in struct {
		name      string
		avatarURL string
	}
	tests := []struct {
		name string
		in   in
		out  User
	}{
		{
			name: "success",
			in: in{
				name:      "name",
				avatarURL: "avatar",
			},
			out: User{
				UpdatedAt: dummyDate,
				Profile: Profile{
					Name:      "name",
					AvatarURL: "avatar",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			patch := monkey.Patch(time.Now, func() time.Time { return dummyDate })
			defer patch.Unpatch()
			u := dummyUser
			u.UpdateProfile(tt.in.name, tt.in.avatarURL)
			assert.Equal(t, tt.out, u)
		})
	}
}

func TestUser_Delete(t *testing.T) {
	dummyDate := time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		in   User
		out  User
		code code.Code
	}{
		{
			name: "failed to already deleted",
			in: User{
				DeletedAt: &dummyDate,
			},
			out: User{
				DeletedAt: &dummyDate,
			},
			code: code.BadRequest,
		},
		{
			name: "success",
			in:   User{},
			out: User{
				DeletedAt: &dummyDate,
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			patch := monkey.Patch(time.Now, func() time.Time { return dummyDate })
			defer patch.Unpatch()
			u := tt.in
			err := u.Delete()
			assert.Equal(t, tt.out, u)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}

func TestUserID_Int(t *testing.T) {
	tests := []struct {
		name string
		in   UserID
		out  int
	}{
		{
			name: "success",
			in:   UserID(1),
			out:  1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := tt.in.Int()
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestAccountID_String(t *testing.T) {
	tests := []struct {
		name string
		in   AccountID
		out  string
	}{
		{
			name: "success",
			in:   AccountID("accid"),
			out:  "accid",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := tt.in.String()
			assert.Equal(t, tt.out, out)
		})
	}
}

func TestEamil_String(t *testing.T) {
	tests := []struct {
		name string
		in   Email
		out  string
	}{
		{
			name: "success",
			in:   Email("mail"),
			out:  "mail",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := tt.in.String()
			assert.Equal(t, tt.out, out)
		})
	}
}
