package usermodel

import "roadmap/pkg/apperror"

var ErrUserNotFound = apperror.New(apperror.ErrUserNotFound, "user not found")
var ErrInvalidPassword = apperror.New(apperror.ErrInvalidPassword, "invalid password")
var ErrUserAlreadyExists = apperror.New(apperror.ErrUserAlreadyExists, "user already exists")
