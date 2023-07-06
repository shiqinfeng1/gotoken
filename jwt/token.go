package jwt

import (
	"time"

	jwtraw "github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	token *jwtraw.Token
	opts  JwtOptions
}

func (j *JwtToken) Generate() (string, error) {
	claims := &jwtraw.RegisteredClaims{
		ExpiresAt: jwtraw.NewNumericDate(time.Now().Add(j.opts.ExpiredByHour * time.Hour)),
		IssuedAt:  jwtraw.NewNumericDate(time.Now()),
		NotBefore: jwtraw.NewNumericDate(time.Now()),
		Issuer:    token.IssuerCA,
		Subject:   "admin",
		ID:        "1",
	}

	j.token = jwtraw.NewWithClaims(jwtraw.SigningMethodHS256, claims)
	return j.token.SignedString(j.opts.Secret)
}
func (j *JwtToken) Refresh() {

}
func (j *JwtToken) Revoke() {

}
func (j *JwtToken) Prase() {

}
func (j *JwtToken) Verify() {

}
