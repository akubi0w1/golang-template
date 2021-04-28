package code

import (
	"net/http"
)

type Code string

const (
	OK              Code = "ok"
	BadRequest      Code = "bad_request"
	InvalidArgument Code = "invalid_argument"
	Unauthorized    Code = "unauthorized"
	NotFound        Code = "not_found"
	Conflict        Code = "conflict"
	JWT             Code = "jwt"
	Session         Code = "session"
	Context         Code = "context"
	Database        Code = "database"
	Password        Code = "password"
	Decode          Code = "decode"
	Encode          Code = "encode"
	UUID            Code = "uuid"
	Storage         Code = "storage"
	Secrets         Code = "secrets"
	CheckImage      Code = "check_image"

	// server error
	Unknown Code = "unknown"
)

func (c Code) ToHTTPStatus() int {
	switch c {
	case OK:
		return http.StatusOK
	case BadRequest:
		return http.StatusBadRequest
	case InvalidArgument:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case NotFound:
		return http.StatusNotFound
	case Conflict:
		return http.StatusConflict
	case JWT:
		return http.StatusInternalServerError
	case Session:
		return http.StatusUnauthorized
	case Context:
		return http.StatusBadRequest
	case Database:
		return http.StatusInternalServerError
	case Password:
		return http.StatusInternalServerError
	case Decode:
		return http.StatusInternalServerError
	case Encode:
		return http.StatusInternalServerError
	case UUID:
		return http.StatusInternalServerError
	case Storage:
		return http.StatusInternalServerError
	case Secrets:
		return http.StatusInternalServerError
	case CheckImage:
		return http.StatusInternalServerError
	case Unknown:
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
