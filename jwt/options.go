package jwt

import "time"

type JwtOptions struct {
	Secret        string
	ExpiredByHour time.Duration
}

var defaultTokenOptions = JwtOptions{
	Secret:        "keep-me-secret",
	ExpiredByHour: 2,
}

type SetOptsFunc func(opts *JwtOptions)

func WithSecret(s string) SetOptsFunc {
	return func(opts *JwtOptions) {
		if s != "" {
			opts.Secret = s
		}
	}
}
func WithExpired(h time.Duration) SetOptsFunc {
	return func(opts *JwtOptions) {
		if h != 0 {
			opts.ExpiredByHour = h
		}
	}
}
