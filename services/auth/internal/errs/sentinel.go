package errs

import "errors"

var (
	ErrorUserNotFound  = errors.New("username not found")
	ErrorPassWrong     = errors.New("password wrong")
	ErrorUsernameExist = errors.New("username exist")
	ErrorDbPg          = errors.New("postgree error")
	ErrorHashBcrypt    = errors.New("hash error")
)
