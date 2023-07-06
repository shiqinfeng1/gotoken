package jwt

import (
	"gotoken"
	"time"

	jwtraw "github.com/golang-jwt/jwt/v5"
)

type builder struct{}

func (b *builder) Build(tOpts gotoken.TokenOptions) gotoken.Token {
	return &JwtTokenSrv{
		accessOpts: jwtOptions{
			secret:        tOpts.Secret,
			signingMethod: tOpts.SigningMethod,
			claims: jwtraw.RegisteredClaims{
				ExpiresAt: jwtraw.NewNumericDate(time.Now().Add(tOpts.ExpiredByHour * time.Hour)),
				IssuedAt:  jwtraw.NewNumericDate(time.Now()),
				NotBefore: jwtraw.NewNumericDate(time.Now()),
				Issuer:    tOpts.Issuer,
				Subject:   tOpts.Subject,
				ID:        tOpts.ID,
				Audience:  tOpts.Audience,
			},
		},
		refreshOpts: jwtOptions{
			secret:        tOpts.Secret,
			signingMethod: tOpts.SigningMethod,
			claims: jwtraw.RegisteredClaims{
				ExpiresAt: jwtraw.NewNumericDate(time.Now().Add(tOpts.RefreshByHour * time.Hour)),
				IssuedAt:  jwtraw.NewNumericDate(time.Now()),
				NotBefore: jwtraw.NewNumericDate(time.Now()),
				Issuer:    tOpts.Issuer,
				Subject:   tOpts.Subject,
				ID:        tOpts.ID,
				Audience:  tOpts.Audience,
			},
		},
	}
}
func NewJwtBuilder() gotoken.TokenBuilder {
	return &builder{}
}
