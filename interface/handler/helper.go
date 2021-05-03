package handler

import (
	"io"

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
