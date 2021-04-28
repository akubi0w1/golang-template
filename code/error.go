package code

import (
	"fmt"

	"golang.org/x/xerrors"
)

type privateError struct {
	code Code
	err  error
}

func (e privateError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.code, e.err)
}

func (e privateError) Unwrap() error {
	return e.err
}

// Errorf formats error with code
func Errorf(c Code, format string, args ...interface{}) error {
	if c == OK {
		return nil
	}
	return privateError{
		code: c,
		err:  xerrors.Errorf(format, args...),
	}
}

// Error format error with code
func Error(c Code, format string) error {
	if c == OK {
		return nil
	}
	return privateError{
		code: c,
		err:  xerrors.New(format),
	}
}

// GetCode gets code
func GetCode(err error) Code {
	if err == nil {
		return OK
	}
	if e, ok := err.(privateError); ok {
		return e.code
	}
	return Unknown
}
