package handler

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/render"
)

type User interface {
	GetAll(w http.ResponseWriter, r *http.Request)
}

type UserImpl struct {
	user usecase.User
}

func NewUser(user usecase.User) *UserImpl {
	return &UserImpl{
		user: user,
	}
}

func (h *UserImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: 実装
	render.Status(r, http.StatusNotImplemented)
	render.JSON(w, r, struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    http.StatusNotImplemented,
		Message: "not implemented",
	})
}
