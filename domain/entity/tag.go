package entity

import "time"

type Tag struct {
	ID        int
	Tag       string
	CreatedAt *time.Time
}

type TagList []Tag
