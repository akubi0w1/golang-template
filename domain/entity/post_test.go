package entity

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/akubi0w1/golang-sample/code"
	"github.com/stretchr/testify/assert"
)

func TestPost_NewPost(t *testing.T) {
	dummyDate := time.Date(2021, 5, 2, 0, 0, 0, 0, time.UTC)
	type in struct {
		title    string
		body     string
		authorID UserID
		tags     TagList
		images   ImageList
	}
	tests := []struct {
		name string
		in   in
		out  Post
		code code.Code
	}{
		{
			name: "success",
			in: in{
				title:    "title001",
				body:     "body001",
				authorID: 1,
				tags: TagList{
					{ID: 1},
				},
				images: ImageList{
					{ID: 1},
				},
			},
			out: Post{
				Title:     "title001",
				Body:      "body001",
				CreatedAt: dummyDate,
				UpdatedAt: dummyDate,
				Tags: TagList{
					{ID: 1},
				},
				Images: ImageList{
					{ID: 1},
				},
				AuthorID: 1,
			},
			code: code.OK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			patch := monkey.Patch(time.Now, func() time.Time { return dummyDate })
			defer patch.Unpatch()
			out, err := NewPost(tt.in.title, tt.in.body, tt.in.authorID, tt.in.tags, tt.in.images)
			assert.Equal(t, tt.out, out)
			assert.Equal(t, tt.code, code.GetCode(err))
		})
	}
}
