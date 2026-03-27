package apperror

type ErrorCode string

const (
	ErrInternalError ErrorCode = "internal_error"
	ErrBadRequest    ErrorCode = "bad_request"
)

// Authentication and Authorization
const (
	ErrAuthInvalidToken ErrorCode = "auth_invalid_token"
	ErrAuthExpiredToken ErrorCode = "auth_expired_token"
)

func (e ErrorCode) String() string {
	return string(e)
}
