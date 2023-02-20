package utils

var (
	CodeClientError   = 1001
	CodeNotFoundError = 1002
	CodeConflictError = 1003
	CodeServerError   = 1004

	CodeCacheMiss = 1005
)

type CustomErrorWrapper struct {
	Message string `json:"message"` // Human readable message for clients
	Code    int    `json:"-"`       // HTTP Status code. We use `-` to skip json marshaling.
	Err     error  `json:"-"`       // The original error. Same reason as above.
}

func (w *CustomErrorWrapper) Error() string {
	// guard against panics
	if w.Err != nil {
		return w.Err.Error()
	}
	return w.Message
}

func NewErrorWrapper(code int, err error, message string) error {
	return &CustomErrorWrapper{
		Message: message,
		Code:    code,
		Err:     err,
	}
}
