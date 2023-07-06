package jwt

import (
	"gotoken"

	jwtraw "github.com/golang-jwt/jwt/v5"
)

type JwtTokenSrv struct {
	accessOpts  jwtOptions
	refreshOpts jwtOptions
}

// jwtraw.SigningMethodHS256
func (j *JwtTokenSrv) Generate() (*gotoken.SignedToken, error) {

	token := &gotoken.SignedToken{}
	var err error

	token.AccessToken, err = func() (string, error) {
		accessToken := jwtraw.NewWithClaims(j.accessOpts.signingMethod, j.accessOpts.claims)
		return accessToken.SignedString(j.accessOpts.secret)
	}()
	if err != nil {
		return nil, err
	}
	token.RefreshToken, err = func() (string, error) {
		refreshToken := jwtraw.NewWithClaims(j.refreshOpts.signingMethod, j.refreshOpts.claims)
		return refreshToken.SignedString(j.refreshOpts.secret)
	}()
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (j *JwtTokenSrv) Refresh() {

}
func (j *JwtTokenSrv) UpdateSecret() {

}
func (j *JwtTokenSrv) Prase() {

}
func (j *JwtTokenSrv) Verify() {

}
