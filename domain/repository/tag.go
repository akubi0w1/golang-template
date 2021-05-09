//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
)

type Tag interface {
	FindByIDs(ctx context.Context, ids []int) (entity.TagList, error)
	FindByTags(ctx context.Context, tags []string) (entity.TagList, error)
	InsertBulk(ctx context.Context, tags entity.TagList) (entity.TagList, error)
	IsExist(ctx context.Context, tag string) (bool, error)
}
