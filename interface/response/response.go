package response

import (
	"net/http"

	"github.com/akubi0w1/golang-sample/code"
	"github.com/akubi0w1/golang-sample/log"
	"github.com/go-chi/render"
)

func NoContent(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

func Success(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, v)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	logger := log.New()

	statusCode := code.GetCode(err).ToHTTPStatus()

	if statusCode == 500 {
		logger.Error("%s", err.Error())
	} else {
		logger.Warn("%s", err.Error())
	}

	msg := err.Error()
	render.Status(r, statusCode)
	render.JSON(w, r, errorResponse{
		Code:    statusCode,
		Message: msg,
	})
}
