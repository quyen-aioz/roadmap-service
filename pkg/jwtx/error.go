package jwtx

import "roadmap/pkg/apperror"

var (
	ErrMissingToken      = apperror.New(apperror.ErrAuthInvalidToken, "missing token")
	ErrInvalidToken      = apperror.New(apperror.ErrAuthInvalidToken, "invalid token")
	ErrMissingSigningKey = apperror.New(apperror.ErrAuthInvalidToken, "missing signing key")
	ErrExpiredToken      = apperror.New(apperror.ErrAuthExpiredToken, "expired token")
)
