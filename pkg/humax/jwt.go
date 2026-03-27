package humax

import (
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

		claims, err := jwtx.GetUserClaims(tokenStr)
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
