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

// TODO: add test
func NewPost(title, body string, authorID UserID, tags TagList, images ImageList) (Post, error) {
	now := time.Now()
	return Post{
		Title:     title,
		Body:      body,
		CreatedAt: now,
		UpdatedAt: now,
		Tags:      tags,
		Images:    images,
		AuthorID:  authorID,
	}, nil
}

type PostList []Post
