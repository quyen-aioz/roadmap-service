package humax

import (
	"context"
	"net/http"
	"roadmap/pkg/jwtx"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

const (
	_bearerPrefix     = "Bearer "
	_claimsContextKey = "claims"
)

var _ Middleware = (*AccessTokenMiddleware)(nil)

type AccessTokenMiddleware struct{}

var (
	accessTokenMiddleware AccessTokenMiddleware
)

func (m *AccessTokenMiddleware) Apply(_ huma.API) func(ctx huma.Context, next func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		tokenStr := ctx.Header("Authorization")

		tokenStr = strings.TrimPrefix(tokenStr, _bearerPrefix)

		if tokenStr == "" {
			ErrorResponse(ctx, http.StatusUnauthorized, jwtx.ErrMissingToken)
			return
		}

		claims, err := jwtx.GetUserClaim(tokenStr)
		if err != nil {
			ErrorResponse(ctx, http.StatusUnauthorized, err)
			return
		}

		nextContext := huma.WithValue(ctx, _claimsContextKey, claims)

		next(nextContext)
	}
}

func RequireAccessToken() Middleware {
	return &accessTokenMiddleware
}

func GetUserIDFromCtx(ctx context.Context) (string, error) {
	claims, ok := ctx.Value(_claimsContextKey).(*jwtx.UserClaim)
	if !ok {
		return "", jwtx.ErrMissingToken
	}

	return claims.UserID, nil
}

func GetUserClaimsFromCtx(ctx context.Context) (*jwtx.UserClaim, error) {
	claims, ok := ctx.Value(_claimsContextKey).(*jwtx.UserClaim)
	if !ok {
		return nil, jwtx.ErrMissingToken
	}

	return claims, nil
}
