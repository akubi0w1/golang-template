package entity

import (
	"time"

	"github.com/akubi0w1/golang-sample/code"
)

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

func NewPost(title, body string, authorID UserID, tags TagList, images ImageList) (Post, error) {
	if err := validateTitle(title); err != nil {
		return Post{}, err
	}
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

func validateTitle(title string) error {
	if title == "" {
		return code.Error(code.BadRequest, "title is required")
	}
	return nil
}

type PostList []Post
