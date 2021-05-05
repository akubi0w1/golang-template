package handler

import (
	"net/http"
	"strconv"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/context"
	"github.com/akubi0w1/golang-sample/interface/request"
	"github.com/akubi0w1/golang-sample/interface/response"
	"github.com/akubi0w1/golang-sample/usecase"
	"github.com/go-chi/chi/v5"
)

type Post interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type PostImpl struct {
	post usecase.Post
}

func NewPost(post usecase.Post) *PostImpl {
	return &PostImpl{
		post: post,
	}
}

func (p *PostImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, total, err := p.post.GetAll(ctx)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toPostListResponse(posts, total))
}

func (p *PostImpl) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
	if err != nil {
		response.Error(w, r, code.Errorf(code.InvalidArgument, "invalid path parameter: %v", err))
		return
	}
	post, err := p.post.GetByID(ctx, postID)
	if err != nil {
		response.Error(w, r, err)
		return
	}
	response.Success(w, r, toPostResponse(post))

}

func (p *PostImpl) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	accountID, err := context.GetAccountID(ctx)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	var req request.CreatePost
	if err := decodeAndValidateRequest(r.Body, &req); err != nil {
		response.Error(w, r, err)
		return
	}

	post, err := p.post.Create(ctx, accountID, req.Title, req.Body, req.Tags, req.ImageIDs)
	if err != nil {
		response.Error(w, r, err)
		return
	}

	response.Success(w, r, toPostResponse(post))
}

func toPostListResponse(ps entity.PostList, total int) response.PostList {
	res := make([]response.Post, 0, len(ps))
	for i := range ps {
		res = append(res, toPostResponse(ps[i]))
	}
	return response.PostList{
		Total: total,
		Posts: res,
	}
}

func toPostResponse(p entity.Post) response.Post {
	return response.Post{
		ID:        p.ID,
		Title:     p.Title,
		Body:      p.Body,
		CreatedAt: dateToString(p.CreatedAt),
		UpdatedAt: dateToString(p.UpdatedAt),
		AuthorID:  p.AuthorID.Int(),
		Tags:      toTagsResponse(p.Tags),
		Images:    toImagesResponse(p.Images),
	}
}

func toTagResponse(t entity.Tag) response.Tag {
	return response.Tag{
		ID:  t.ID,
		Tag: t.Tag,
	}
}

func toTagsResponse(ts entity.TagList) []response.Tag {
	tags := make([]response.Tag, 0, len(ts))
	for i := range ts {
		tags = append(tags, toTagResponse(ts[i]))
	}
	return tags
}

func toImageResponse(image entity.Image) response.Image {
	return response.Image{
		ID:  image.ID,
		URL: image.URL,
	}
}

func toImagesResponse(imgs entity.ImageList) []response.Image {
	images := make([]response.Image, 0, len(imgs))
	for i := range imgs {
		images = append(images, toImageResponse(imgs[i]))
	}
	return images
}
