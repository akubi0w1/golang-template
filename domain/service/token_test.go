package service

import (
	"testing"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	mock_repository "github.com/akubi0w1/golang-sample/mock/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTokenManager_Generate(t *testing.T) {
	tests := []struct {
		name     string
		injector func(mt *mock_repository.MockJWT)
		in       entity.AccountID
		out      entity.Token
		code     code.Code
	}{
		{
			name:     "failed to accountID is empty",
			injector: func(mt *mock_repository.MockJWT) {},
			in:       "",
			out:      "",
			code:     code.Session,
		},
		{
			name: "failed to generate token",
			injector: func(mt *mock_repository.MockJWT) {
				mt.EXPECT().Generate(entity.Claims{AccountID: "accid"}).Return(entity.Token(""), code.Error(code.JWT, "some error"))
			},
			in:   "accid",
			out:  "",
			code: code.JWT,
		},
		{
			name: "success",
			injector: func(mt *mock_repository.MockJWT) {
				mt.EXPECT().Generate(entity.Claims{AccountID: "accid"}).Return(entity.Token("token001"), nil)
			},
			in:   "accid",
			out:  entity.Token("token001"),
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mt := mock_repository.NewMockJWT(ctrl)
			tt.injector(mt)
			srv := NewTokenManager(mt)
			out, err := srv.Generate(tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))

		})
	}
}
