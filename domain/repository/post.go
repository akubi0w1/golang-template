//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
)

type Post interface {
	Insert(ctx context.Context, post entity.Post) (entity.Post, error)
	FindAll(ctx context.Context) (entity.PostList, error)
	FindByID(ctx context.Context, id int) (entity.Post, error)
	Count(ctx context.Context) (int, error)
}
