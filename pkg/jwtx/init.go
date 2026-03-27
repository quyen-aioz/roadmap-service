package jwtx

var (
	_signingKey string
)

func InitJWT(signingKey string) error {
	if len(signingKey) == 0 {
		return ErrMissingSigningKey
	}

	_signingKey = signingKey

	return nil
}
