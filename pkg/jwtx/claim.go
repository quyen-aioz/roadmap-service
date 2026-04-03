package jwtx

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GetUserClaim(tokenString string) (*UserClaim, error) {
	if len(_signingKey) == 0 {
		return nil, ErrMissingSigningKey
	}

	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(_signingKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken.Wrap(err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		return nil, ErrInvalidToken.Wrap(errors.New("invalid token claims"))
	}

	return claims, nil
}

func GenerateToken(claims UserClaim, duration time.Duration) (string, error) {
	if len(_signingKey) == 0 {
		return "", ErrMissingSigningKey
	}

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(_signingKey))
}
