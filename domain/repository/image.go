//go:generate mockgen -source=$GOFILE -destination=../../mock/$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE -build_flags=-mod=mod
package repository

import (
	"context"

	"github.com/akubi0w1/golang-sample/domain/entity"
)

type Image interface {
	FindByIDs(ctx context.Context, ids []int) (entity.ImageList, error)
}
