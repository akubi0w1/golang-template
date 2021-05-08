//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package service

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
	"github.com/akubi0w1/golang-sample/domain/repository"
)

type Tag interface {
	CreateOrGetMultiple(ctx context.Context, tags []string) (entity.TagList, error)
}

type TagImpl struct {
	tag repository.Tag
}

func NewTag(tag repository.Tag) *TagImpl {
	return &TagImpl{
		tag: tag,
	}
}

func (t *TagImpl) CreateOrGetMultiple(ctx context.Context, tags []string) (entity.TagList, error) {
	if len(tags) == 0 {
		return entity.TagList{}, nil
	}

	newTags := entity.TagList{}
	for i := range tags {
		isExist, err := t.tag.IsExist(ctx, tags[i])
		if err != nil {
			return entity.TagList{}, err
		}
		if !isExist {
			_newTag, err := entity.NewTag(tags[i])
			if err != nil {
				return entity.TagList{}, err
			}
			newTags = append(newTags, _newTag)
		}
	}
	if len(newTags) > 0 {
		_, err := t.tag.InsertBulk(ctx, newTags)
		if err != nil {
			return entity.TagList{}, err
		}
	}
	tagList, err := t.tag.FindByTags(ctx, tags)
	if err != nil {
		return entity.TagList{}, err
	}
	return tagList, nil
}
