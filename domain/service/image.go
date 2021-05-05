package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type Image interface {
	GetByIDs(ctx context.Context, ids []int) (entity.ImageList, error)
}

type ImageImpl struct {
	image repository.Image
}

func NewImage(image repository.Image) *ImageImpl {
	return &ImageImpl{
		image: image,
	}
}

// TODO: add test
func (ii *ImageImpl) GetByIDs(ctx context.Context, ids []int) (entity.ImageList, error) {
	return ii.image.FindByIDs(ctx, ids)
}
