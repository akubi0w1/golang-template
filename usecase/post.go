//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package usecase

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/service"
)

type Post interface {
	Create(ctx context.Context, accountID entity.AccountID, title, body string, tags []string, imageIDs []int) (entity.Post, error)
	GetAll(ctx context.Context) (posts entity.PostList, total int, err error)
	GetByID(ctx context.Context, id int) (entity.Post, error)
}

type PostImpl struct {
	post  service.Post
	user  service.User
	tag   service.Tag
	image service.Image
}

func NewPost(post service.Post, user service.User, tag service.Tag, image service.Image) *PostImpl {
	return &PostImpl{
		post:  post,
		user:  user,
		tag:   tag,
		image: image,
	}
}

func (p *PostImpl) Create(ctx context.Context, accountID entity.AccountID, title, body string, tags []string, imageIDs []int) (entity.Post, error) {
	author, err := p.user.GetByAccountID(ctx, accountID)
	if err != nil {
		return entity.Post{}, err
	}

	tagList, err := p.tag.CreateOrGetMultiple(ctx, tags)
	if err != nil {
		return entity.Post{}, err
	}
	imageList, err := p.image.GetByIDs(ctx, imageIDs)
	if err != nil {
		return entity.Post{}, err
	}

	post, err := p.post.Create(ctx, title, body, author.ID, tagList, imageList)
	if err != nil {
		return entity.Post{}, err
	}
	return post, nil
}

func (p *PostImpl) GetAll(ctx context.Context) (entity.PostList, int, error) {
	return p.post.GetAll(ctx)
}

func (p *PostImpl) GetByID(ctx context.Context, id int) (entity.Post, error) {
	return p.post.GetByID(ctx, id)
}
