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

// Roadmap
const (
	ErrInvalidStatus ErrorCode = "invalid_status"
	ErrNotFound      ErrorCode = "not_found"
)

func (e ErrorCode) String() string {
	return string(e)
}
