package repositories

import "runtime"

// adapted from here: https://dev.to/tigorlazuardi/go-creating-custom-error-wrapper-and-do-proper-error-equality-check-11k7
var (
	CodeClientError   = 1001
	CodeNotFoundError = 1002
	CodeConflictError = 1003
	CodeServerError   = 1004

	CodeCacheMiss = 1005
)

var (
	MsgServerError               = "server error"
	MsgClientBadFormattedRequest = "bad format request"
	MsgClientNotFoundRequest     = "resource not found"
)

type ErrorWrapper struct {
	Message    string `json:"message"` // human readable error
	Code       int    `json:"-"`       // code
	Err        error  `json:"-"`       // original error
	Filename   string `json:"-"`
	LineNumber int    `json:"-"`
}

func (w *ErrorWrapper) Error() string {
	// guard against panics
	if w.Err != nil {
		return w.Err.Error()
	}
	return w.Message
}

func NewErrorWrapper(code int, msg string, err error) *ErrorWrapper {
	// getting previous call stack file and line info
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       code,
		Message:    msg,
		Err:        err,
		Filename:   filename,
		LineNumber: line,
	}
}

func NewBadFormattedRequest() *ErrorWrapper {
	return NewErrorWrapper(CodeClientError, MsgClientBadFormattedRequest, nil)
}
