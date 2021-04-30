package entity

import "time"

type Post struct {
	ID        int
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Author    User
	Tags      TagList
	Images    ImageList
}

type PostList []Post
