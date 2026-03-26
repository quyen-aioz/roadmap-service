package apperror

type ErrorCode string

const (
	ErrInternalError ErrorCode = "internal_error"
	ErrBadRequest    ErrorCode = "bad_request"
)

func (e ErrorCode) String() string {
	return string(e)
}
