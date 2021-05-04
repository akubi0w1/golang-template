package entity

import "time"

type Tag struct {
	ID        int
	Tag       string
	CreatedAt time.Time
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
