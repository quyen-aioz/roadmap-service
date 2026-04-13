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
	ErrInvalidStatus        ErrorCode = "invalid_status"
	ErrInvalidGroupID       ErrorCode = "invalid_group"
	ErrRoadmapNotFound      ErrorCode = "group_not_found"
	ErrUnsupportedMediaType ErrorCode = "unsupported_media_type"
)

// User
const (
	ErrUserNotFound      ErrorCode = "user_not_found"
	ErrInvalidPassword   ErrorCode = "invalid_password"
	ErrUserAlreadyExists ErrorCode = "user_already_exists"
)

func (e ErrorCode) String() string {
	return string(e)
}
