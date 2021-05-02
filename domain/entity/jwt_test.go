package entity

import (
	"testing"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/stretchr/testify/assert"
)

func TestToken_String(t *testing.T) {
	tests := []struct {
		name string
		in   Token
		out  string
	}{
		{
			name: "success",
			in:   Token("token"),
			out:  "token",
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

func TestClaims_NewClaims(t *testing.T) {
	tests := []struct {
		name string
		in   AccountID
		out  Claims
		code code.Code
	}{
		{
			name: "failed to accountID is empty",
			in:   AccountID(""),
			out:  Claims{},
			code: code.Session,
		},
		{
			name: "success",
			in:   AccountID("accid"),
			out: Claims{
				AccountID: "accid",
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out, err := NewClaims(tt.in)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
