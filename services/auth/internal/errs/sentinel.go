package errs

import "errors"

var (
	ErrorUserNotFound = errors.New("username not found")
	ErrorPassWrong    = errors.New("password wrong")
)
