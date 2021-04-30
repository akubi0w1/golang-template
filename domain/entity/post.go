package entity

import "time"

type Post struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Tags      TagList
	Images    ImageList
	AuthorID  UserID
}

type PostList []Post
