//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type Post interface {
	Create(ctx context.Context, title, body string, authorID entity.UserID, tags entity.TagList, images entity.ImageList) (entity.Post, error)
	GetAll(ctx context.Context) (posts entity.PostList, total int, err error)
	GetByID(ctx context.Context, id int) (entity.Post, error)
}

type PostImpl struct {
	post repository.Post
}

func NewPost(post repository.Post) *PostImpl {
	return &PostImpl{
		post: post,
	}
}

// TODO: add test
func (p *PostImpl) Create(ctx context.Context, title, body string, authorID entity.UserID, tags entity.TagList, images entity.ImageList) (entity.Post, error) {
	post, err := entity.NewPost(title, body, authorID, tags, images)
	if err != nil {
		return entity.Post{}, err
	}
	post, err = p.post.Insert(ctx, post)
	if err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

// TODO: add test
func (p *PostImpl) GetAll(ctx context.Context) (entity.PostList, int, error) {
	posts, err := p.post.FindAll(ctx)
	if err != nil {
		return entity.PostList{}, 0, err
	}
	total, err := p.post.Count(ctx)
	if err != nil {
		return entity.PostList{}, 0, err
	}
	return posts, total, nil
}

// TODO: add test
func (p *PostImpl) GetByID(ctx context.Context, id int) (entity.Post, error) {
	return p.post.FindByID(ctx, id)
}
