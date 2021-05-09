package handler

import (
	"io"
	"time"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/interface/request"
	"github.com/go-chi/render"
)

func decodeAndValidateRequest(body io.ReadCloser, req request.RequestType) error {
	if err := render.DecodeJSON(body, req); err != nil {
		return code.Errorf(code.Decode, "failed to decode request body: %v", err)
	}
	if err := req.Validate(); err != nil {
		return err
	}
	return nil
}

func dateToString(t time.Time) string {
	return t.Format("2006/01/02 03:04:05")
}
