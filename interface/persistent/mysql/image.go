package mysql

import (
	"context"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	entimage "github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/image"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent/predicate"
)

type ImageImpl struct {
	cli *ent.Client
}

func NewImage(cli *ent.Client) *ImageImpl {
	return &ImageImpl{
		cli: cli,
	}
}

func (ii *ImageImpl) FindByIDs(ctx context.Context, ids []int) (entity.ImageList, error) {
	if len(ids) == 0 {
		return entity.ImageList{}, nil
	}
	conds := make([]predicate.Image, 0, len(ids))
	for i := range ids {
		conds = append(conds, entimage.IDEQ(ids[i]))
	}
	images, err := ii.cli.Image.Query().Where(entimage.Or(conds...)).WithUploadedBy().All(ctx)
	if err != nil {
		return entity.ImageList{}, code.Errorf(code.Database, "failed to find images by ids: %v", err)
	}
	return toEntityImageList(images), nil
}

func toEntityImage(image *ent.Image) entity.Image {
	return entity.Image{
		ID:        image.ID,
		URL:       image.URL,
		CreatedBy: entity.UserID(image.Edges.UploadedBy.ID),
	}
}

func toEntityImageList(images []*ent.Image) entity.ImageList {
	list := make(entity.ImageList, 0, len(images))
	for i := range images {
		list = append(list, toEntityImage(images[i]))
	}
	return list
}
