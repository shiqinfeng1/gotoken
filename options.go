package gotoken

import (
	"time"

	jwtraw "github.com/golang-jwt/jwt/v5"
)

type SigningMethod jwtraw.SigningMethod

var (
	JwtSignES256 SigningMethod = jwtraw.SigningMethodES256
	JwtSignES384 SigningMethod = jwtraw.SigningMethodES384
	JwtSignES512 SigningMethod = jwtraw.SigningMethodES512
	JwtSignEdDSA SigningMethod = jwtraw.SigningMethodEdDSA
	JwtSignRS256 SigningMethod = jwtraw.SigningMethodRS256
	JwtSignRS384 SigningMethod = jwtraw.SigningMethodRS384
	JwtSignRS512 SigningMethod = jwtraw.SigningMethodRS512
	JwtSignHS256 SigningMethod = jwtraw.SigningMethodHS256
	JwtSignHS384 SigningMethod = jwtraw.SigningMethodHS384
	JwtSignHS512 SigningMethod = jwtraw.SigningMethodHS512
)

type TokenOptions struct {
	Secret        any
	SigningMethod SigningMethod
	ExpiredByHour time.Duration
	RefreshByHour time.Duration
	Issuer        string
	Subject       string
	Audience      []string
	ID            string
}

var defaultTokenOptions = TokenOptions{
	Secret:        []byte("keep-me-secret"),
	ExpiredByHour: 2,
	RefreshByHour: 30 * 24,
	SigningMethod: JwtSignHS256,
	Issuer:        "",
	Subject:       "",
	Audience:      []string{},
	ID:            "",
}

type SetOptsFunc func(opts *TokenOptions)

func WithSecret(s any) SetOptsFunc {
	return func(opts *TokenOptions) {
		if s != nil {
			opts.Secret = s
		}
	}
}
func WithExpired(h time.Duration) SetOptsFunc {
	return func(opts *TokenOptions) {
		if h != 0 {
			opts.ExpiredByHour = h
		}
	}
}
func WithRefresh(h time.Duration) SetOptsFunc {
	return func(opts *TokenOptions) {
		if h != 0 {
			opts.RefreshByHour = h
		}
	}
}
func WithIssuer(issuer string) SetOptsFunc {
	return func(opts *TokenOptions) {
		if issuer != "" {
			opts.Issuer = issuer
		}
	}
}

func WithSubject(s string) SetOptsFunc {
	return func(opts *TokenOptions) {
		if s != "" {
			opts.Subject = s
		}
	}
}

func WithID(id string) SetOptsFunc {
	return func(opts *TokenOptions) {
		if id != "" {
			opts.ID = id
		}
	}
}

func WithAudience(ads []string) SetOptsFunc {
	return func(opts *TokenOptions) {
		if len(ads) != 0 {
			opts.Audience = ads
		}
	}
}
func WithSignMethod(sm SigningMethod) SetOptsFunc {
	return func(opts *TokenOptions) {
		if sm != nil {
			opts.SigningMethod = sm
		}
	}
}
