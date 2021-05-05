package entity

import (
	"time"

	"github.com/akubi0w1/golang-sample/code"
)

type Tag struct {
	ID        int
	Tag       string
	CreatedAt time.Time
}

func NewTag(tag string) (Tag, error) {
	if tag == "" {
		return Tag{}, code.Errorf(code.BadRequest, "tag is empty")
	}
	now := time.Now()
	return Tag{
		Tag:       tag,
		CreatedAt: now,
	}, nil
}

type TagList []Tag

// TODO: add test
func (tl TagList) GetIDs() []int {
	ids := make([]int, 0, len(tl))
	for i := range tl {
		ids = append(ids, tl[i].ID)
	}
	return ids
}
