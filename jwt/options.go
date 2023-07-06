package jwt

import jwtraw "github.com/golang-jwt/jwt/v5"

type jwtOptions struct {
	secret        any
	signingMethod jwtraw.SigningMethod
	claims        jwtraw.RegisteredClaims
}
