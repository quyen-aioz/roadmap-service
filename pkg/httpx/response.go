package httpx

import (
	"fmt"
	"roadmap/pkg/apperror"
)

var DebugMsgEnabled = false // false in prod

type Response = ResponseT[any]

type ResponseT[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`

	DebugMsg        string            `json:"debug,omitempty"`
	ValidationError []ValidationError `json:"validationError,omitempty"`
}

type Paging[T any] struct {
	Total int64 `json:"total"`
	Data  []T   `json:"data"`
}

type ValidationError struct {
	Field string `json:"field"`
	Param string `json:"param,omitempty"`
	Tag   string `json:"tag"`
}

func ErrorResponseT[T any](err error) ResponseT[T] {
	appErr := apperror.GetAppError(err)

	resp := ResponseT[T]{
		Success: false,
		Code:    appErr.ErrorCode(),
		Message: appErr.ErrorMessage(),
	}

	// debug message only enabled in development environment
	if DebugMsgEnabled {
		resp.DebugMsg = fmt.Sprintf("Err: %+v", err)
	}

	return resp
}

func SuccessResponseT[T any](data T) ResponseT[T] {
	return ResponseT[T]{
		Success: true,
		Code:    "",
		Data:    data,
	}
}

func AutoResponseT[T any](data T, err error) ResponseT[T] {
	if err != nil {
		return ErrorResponseT[T](err)
	}

	return SuccessResponseT(data)
}
