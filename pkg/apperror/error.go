package apperror

import (
	"errors"
	"fmt"
)

const (
	DefaultMessage = "An internal error has occurred. Please contact technical support."
)

type AppErrorer interface {
	error
	ErrorCode() string
	ErrorMessage() string
	Unwrap() error
	Wrap(err error) AppErrorer
}

type appError struct {
	code    ErrorCode
	message string
	cause   error
}

func (e *appError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.code, e.message, e.cause)
	}
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

func (e *appError) ErrorCode() string    { return e.code.String() }
func (e *appError) ErrorMessage() string { return e.message }
func (e *appError) Unwrap() error        { return e.cause }
func (e *appError) Wrap(err error) AppErrorer {
	return &appError{
		code:    e.code,
		message: e.message,
		cause:   errors.Join(e.cause, err),
	}
}

func New(code ErrorCode, message string) AppErrorer {
	return &appError{
		code:    code,
		message: message,
	}
}

func Wrap(err error, code ErrorCode, message string) AppErrorer {
	return &appError{
		code:    code,
		message: message,
		cause:   err,
	}
}

func GetAppError(err error) AppErrorer {
	var appErr AppErrorer
	if errors.As(err, &appErr) {
		return appErr
	}

	return Wrap(err, ErrInternalError, DefaultMessage)
}
