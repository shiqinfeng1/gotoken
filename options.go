package gotoken

import "time"

type TokenOptions struct {
	Secret        string
	ExpiredByHour time.Duration
}

var defaultTokenOptions = TokenOptions{
	Secret:        "keep-me-secret",
	ExpiredByHour: 2,
}

type SetOptsFunc func(opts *TokenOptions)

func WithSecret(s string) SetOptsFunc {
	return func(opts *TokenOptions) {
		if s != "" {
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
