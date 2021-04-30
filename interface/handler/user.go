package handler

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/request"
	"github.com/akubi0w1/golang-sample/interface/response"
	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/render"
)

type User interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
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

func (h *UserImpl) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req *request.CreateUser
	if err := render.DecodeJSON(r.Body, &req); err != nil {
		response.Error(w, r, code.Errorf(code.Decode, "failed to decode request body: %v", err))
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, r, err)
		return
	}

	user, err := h.user.CreateWithProfile(ctx, entity.AccountID(req.AccountID), entity.Email(req.Email), req.Password, req.Name, req.AvatarURL)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toUserResponse(user))
}

func toUserResponse(user entity.User) response.User {
	return response.User{
		ID:        user.ID,
		AccountID: user.AccountID.String(),
		Email:     user.Email.String(),
		Profile: response.Profile{
			ID:        user.Profile.ID,
			Name:      user.Profile.Name,
			AvatarURL: user.Profile.AvatarURL,
		},
	}
}
