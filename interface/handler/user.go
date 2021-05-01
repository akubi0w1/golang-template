package handler

import (
	"net/http"
	"strconv"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/context"
	"github.com/akubi0w1/golang-sample/interface/request"
	"github.com/akubi0w1/golang-sample/interface/response"
	"github.com/akubi0w1/golang-sample/interface/session"
	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type User interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Authorize(w http.ResponseWriter, r *http.Request)
}

type UserImpl struct {
	user usecase.User
}

func NewUser(user usecase.User) *UserImpl {
	return &UserImpl{
		user: user,
	}
}

// TODO: add test
func (h *UserImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, total, err := h.user.GetAll(ctx, entity.WithLimit(10))
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toUserListResponse(users, total))
}

// TODO: add test
func (h *UserImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		response.Error(w, r, code.Errorf(code.InvalidArgument, "invalid path parameter: %v", err))
		return
	}

	user, err := h.user.GetByID(ctx, entity.UserID(userID))
	if err != nil {
		response.Error(w, r, err)
		return
	}
	response.Success(w, r, toUserResponse(user))
}

// TODO: add test
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

// TODO: add test
func (h *UserImpl) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountID, err := context.GetAccountID(ctx)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	var req request.UpdateUser
	if err := decodeAndValidateRequest(r.Body, &req); err != nil {
		response.Error(w, r, err)
		return
	}

	user, err := h.user.UpdateProfile(ctx, accountID, req.Name, req.AvatarURL)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toUserResponse(user))
}

// TODO: add test
func (h *UserImpl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountID, err := context.GetAccountID(ctx)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	err = h.user.Delete(ctx, accountID)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.NoContent(w, r)
}

// TODO: add test
func (h *UserImpl) Authorize(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req request.Login
	if err := decodeAndValidateRequest(r.Body, &req); err != nil {
		response.Error(w, r, err)
		return
	}

	_, token, err := h.user.Authorize(ctx, entity.AccountID(req.AccountID), req.Password)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	session.Start(w, token)
	response.Success(w, r, toTokenResponse(token))
}

func toUserResponse(user entity.User) response.User {
	return response.User{
		ID:        user.ID.Int(),
		AccountID: user.AccountID.String(),
		Email:     user.Email.String(),
		Profile: response.Profile{
			ID:        user.Profile.ID,
			Name:      user.Profile.Name,
			AvatarURL: user.Profile.AvatarURL,
		},
	}
}

func toUserListResponse(users entity.UserList, total int) response.UserList {
	res := make([]response.User, 0, len(users))
	for i := range users {
		res = append(res, toUserResponse(users[i]))
	}
	return response.UserList{
		Total: total,
		Users: res,
	}
}

func toTokenResponse(token entity.Token) response.Token {
	return response.Token{
		Token: token.String(),
	}
}
