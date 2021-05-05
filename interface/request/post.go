package request

import "github.com/akubi0w1/golang-sample/code"

type CreatePost struct {
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Tags     []string `json:"tags"`
	ImageIDs []int    `json:"imageIds"`
}

func (req *CreatePost) Validate() error {
	if req.Title == "" || req.Body == "" {
		return code.Error(code.InvalidArgument, "require fields is empty")
	}
	return nil
}
