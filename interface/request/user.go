package request

import "github.com/akubi0w1/golang-sample/code"

type CreateUser struct {
	AccountID       string `json:"accountId"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	AvatarURL       string `json:"avatarUrl"`
}

func (req *CreateUser) Validate() error {
	if req.AccountID == "" || req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return code.Error(code.InvalidArgument, "require fields is empty")
	}
	if req.Password != req.PasswordConfirm {
		return code.Error(code.BadRequest, "password not equal password confirm")

	}
	return nil
}

type Login struct {
	AccountID string `json:"accountId"`
	Password  string `json:"password"`
}

func (req *Login) Validate() error {
	if req.AccountID == "" || req.Password == "" {
		return code.Error(code.InvalidArgument, "reequire fields is empty")
	}
	return nil
}

type UpdateUser struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
}

func (req *UpdateUser) Validate() error {
	return nil
}
